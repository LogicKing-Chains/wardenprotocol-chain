﻿---
sidebar_position: 4
---

# External modules

## Overview

In this section, you'll find descriptions of external [Cosmos SDK](https://docs.cosmos.network) modules used by the Warden Protocol. For more details, follow the provided links to external documentation.

## x/evm

*This module is coming soon.*

The `x/evm` module is a Cosmos SDK module by **Evmos** that allows for the deployment of smart contracts, interaction with the Ethereum Virtual Machine (EVM), and the use of EVM tooling.

To learn more, see the following:

- [`x/evm` in Evmos documentation](https://docs.evmos.org/protocol/modules/evm)
- [`x/evm` on GitHub](https://github.com/evmos/ethermint/tree/v0.22.0/x/evm)

## x/gmp

The `x/gmp` module is a Cosmos SDK module by **Axelar** that enables Axelar General Message Passing: sending and receiving messages on EVM and Cosmos chains.

To learn more, see the following:

- [Cosmos GMP documentation](https://docs.axelar.dev/dev/cosmos-gmp)
- [Axelar GMP SDK on GitHub](https://github.com/axelarnetwork/axelar-gmp-sdk-solidity)

<!---
We should add more details here since there are things that are unique for our chain.
--->

## x/oracle

*This module is coming soon.*

The `x/oracle` module is a Cosmos SDK module by **Skip Protocol** that enables storing prices on-chain in **Slinky** (an [oracle service](/learn/glossary#oracle-service)).

To learn more, see the following:

- [Slinky documentation](https://docs.skip.money/slinky/overview)
- [`x/oracle` on GitHub](https://github.com/skip-mev/slinky/tree/main/x/oracle)

## x/wasm

The `x/wasm` module is a Cosmos SDK module that processes certain messages and uses them to upload, instantiate, and execute smart contracts. It's an integral part of **CosmWasm** – a smart contract platform that can be integrated into any blockchain built on top of the Cosmos SDK.

To learn more, see the following:

- [CosmWasm Contract Semantics](https://docs.cosmwasm.com/docs/smart-contracts/contract-semantics)
- [`x/wasm` on GitHub](https://github.com/CosmWasm/wasmd/blob/main/x/wasm)