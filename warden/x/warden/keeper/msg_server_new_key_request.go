package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/warden-protocol/wardenprotocol/warden/x/warden/types/v1beta3"
)

func (k msgServer) NewKeyRequest(ctx context.Context, msg *types.MsgNewKeyRequest) (*types.MsgNewKeyRequestResponse, error) {
	if err := k.assertActAuthority(msg.Authority); err != nil {
		return nil, err
	}

	creator := k.actKeeper.GetActionCreator(ctx)

	if _, err := k.SpacesKeeper.Get(ctx, msg.SpaceId); err != nil {
		return nil, err
	}

	keychain, err := k.keychains.Get(ctx, msg.KeychainId)
	if err != nil {
		return nil, err
	}

	if keychain.Fees != nil {
		err2 := chargeKeychainFee(&k, ctx, msg.MaxFees, keychain.AccAddress(), keychain.Fees.KeyReq, creator)
		if err2 != nil {
			return nil, err2
		}
	}

	req := &types.KeyRequest{
		Creator:    creator,
		SpaceId:    msg.SpaceId,
		KeychainId: msg.KeychainId,
		KeyType:    msg.KeyType,
		Status:     types.KeyRequestStatus_KEY_REQUEST_STATUS_PENDING,
		RuleId:     msg.RuleId,
	}

	id, err := k.keyRequests.Append(ctx, req)
	if err != nil {
		return nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if err := sdkCtx.EventManager().EmitTypedEvent(&types.EventNewKeyRequest{
		Id:         id,
		SpaceId:    req.SpaceId,
		KeychainId: req.KeychainId,
		RuleId:     req.RuleId,
		KeyType:    req.KeyType,
		Creator:    req.Creator,
	}); err != nil {
		return nil, err
	}

	return &types.MsgNewKeyRequestResponse{
		Id: id,
	}, nil
}
