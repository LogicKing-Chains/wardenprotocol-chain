import { CONTRACTS } from '@certusone/wormhole-sdk';
import {
  deriveAddress,
  getTokenBridgeDerivedAccounts,
  getTransferNativeWithPayloadCpiAccounts,
} from '@certusone/wormhole-sdk/lib/cjs/solana/index.js';
import { deriveEndpointKey } from '@certusone/wormhole-sdk/lib/cjs/solana/tokenBridge/index.js';
import { getProgramSequenceTracker } from '@certusone/wormhole-sdk/lib/cjs/solana/wormhole/index.js';
import { Idl, Program } from '@coral-xyz/anchor';
import { getAssociatedTokenAddressSync } from '@solana/spl-token';
import {
  Connection,
  Keypair,
  PublicKey,
  PublicKeyInitData,
  Transaction,
  sendAndConfirmTransaction,
} from '@solana/web3.js';
import { ChainId } from '@wormhole-foundation/sdk';
import bs58 from 'bs58';

import { GmpWithToken, IDL as IdlGmpWithToken } from '../../contracts/solana/target/types/gmp_with_token.js';
import { config } from '../config/schema.js';
import { getWormholeContractsNetwork } from '../utils.js';

export class SolanaGmpWithTokenClient {
  wormholeCore: string;
  wormholeTokenBridge: string;
  programId: PublicKey;
  adminKeypair: Keypair;
  connection: Connection;

  constructor() {
    this.wormholeCore = CONTRACTS[getWormholeContractsNetwork(config.ENVIRONMENT)].solana.core;
    this.wormholeTokenBridge = CONTRACTS[getWormholeContractsNetwork(config.ENVIRONMENT)].solana.token_bridge;
    this.programId = new PublicKey(config.SOLANA_GMP_WITH_TOKEN_CONTRACT_ADDRESS);
    this.adminKeypair = Keypair.fromSecretKey(bs58.decode(config.SOL_PRIVATE_KEY));
  }

  async initialize(): Promise<void> {
    const program = createProgram<GmpWithToken>(this.programId, IdlGmpWithToken, this.adminKeypair.publicKey);

    const instruction = await program.methods
      .initialize(config.SOLANA_GMP_WITH_TOKEN_RELAYER_FEE, config.SOLANA_GMP_WITH_TOKEN_RELAYER_FEE_PRECISION)
      .accounts({
        owner: this.adminKeypair.publicKey,
        senderConfig: deriveAddress([Buffer.from('sender')], this.programId),
        redeemerConfig: deriveAddress([Buffer.from('redeemer')], this.programId),
        tokenBridgeProgram: new PublicKey(this.wormholeTokenBridge),
        wormholeProgram: new PublicKey(this.wormholeCore),
        ...getTokenBridgeDerivedAccounts(this.programId, this.wormholeTokenBridge, this.wormholeCore),
      })
      .instruction();

    const tx = new Transaction().add(instruction);
    const result = await sendAndConfirmTransaction(this.connection, tx, [this.adminKeypair]);

    console.log(result);
  }

  async registerForeignContract(
    foreignChain: ChainId,
    foreignContract: Buffer,
    foreignTokenBridge: string,
  ): Promise<void> {
    const program = createProgram<GmpWithToken>(this.programId, IdlGmpWithToken, this.adminKeypair.publicKey);
    // const test = new PublicKey('0x27b6Fa47efd7Eb3F67ED4A28703EC907A96C2f97').toBuffer();

    const foreignContractKey = deriveAddress(
      [
        Buffer.from('foreign_contract'),
        (() => {
          const buf = Buffer.alloc(2);
          buf.writeUInt16LE(foreignChain);
          return buf;
        })(),
      ],
      this.programId,
    );

    const instruction = await program.methods
      .registerForeignContract(foreignChain, [...foreignContract])
      .accounts({
        owner: this.adminKeypair.publicKey,
        config: deriveAddress([Buffer.from('sender')], this.programId),
        foreignContract: foreignContractKey,
        tokenBridgeForeignEndpoint: deriveEndpointKey(this.wormholeTokenBridge, foreignChain, foreignTokenBridge),
        tokenBridgeProgram: new PublicKey(this.wormholeTokenBridge),
      })
      .instruction();

    const tx = new Transaction().add(instruction);
    const result = await sendAndConfirmTransaction(this.connection, tx, [this.adminKeypair]);

    console.log(result);
  }

  async sendWrapped(from: Keypair, foreignChain: ChainId, mint: PublicKey): Promise<void> {
    const program = createProgram<GmpWithToken>(this.programId, IdlGmpWithToken);
    const tracker = await getProgramSequenceTracker(this.connection, this.wormholeTokenBridge, this.wormholeCore);
    const message = deriveAddress(
      [
        Buffer.from('bridged'),
        (() => {
          const buf = Buffer.alloc(8);
          buf.writeBigUInt64LE(tracker.sequence + 1n);
          return buf;
        })(),
      ],
      this.programId,
    );

    const fromTokenAccount = getAssociatedTokenAddressSync(mint, from.publicKey);
    const tmpTokenAccount = deriveAddress([Buffer.from('tmp'), new PublicKey(mint).toBuffer()], this.programId);

    const tokenBridgeAccounts = getTransferNativeWithPayloadCpiAccounts(
      this.programId,
      this.wormholeTokenBridge,
      this.wormholeCore,
      from,
      message,
      fromTokenAccount,
      mint,
    );

    // const instruction = this.program.methods
    //   .sendNativeTokensWithPayload(
    //     params.batchId,
    //     new BN(params.amount.toString()),
    //     [...params.recipientAddress],
    //     params.recipientChain,
    //   )
    //   .accounts({
    //     config: deriveSenderConfigKey(programId),
    //     foreignContract: deriveForeignContractKey(programId, params.recipientChain),
    //     tmpTokenAccount,
    //     tokenBridgeProgram: new PublicKey(tokenBridgeProgramId),
    //     ...tokenBridgeAccounts,
    //   })
    //   .instruction();
  }
}

function createProgram<IDL extends Idl = Idl>(programId: PublicKeyInitData, idl: IDL, payer?: PublicKeyInitData) {
  const connection = new Connection(config.SOLANA_RPC, 'processed');

  return new Program<IDL>(idl, new PublicKey(programId), {
    connection,
    publicKey: payer == undefined ? undefined : new PublicKey(payer),
  });
}
