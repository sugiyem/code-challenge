package keeper

import (
	"context"
	"fmt"

	"crude/x/crude/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateResource(goCtx context.Context, msg *types.MsgUpdateResource) (*types.MsgUpdateResourceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var resource = types.Resource{
		Metadata: msg.Metadata,
		Value:    msg.Value,
		Creator:  msg.Creator,
		Id:       msg.Id,
	}

	val, found := k.GetResource(ctx, msg.Id)
	if !found {
		// there is no resource with given id
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("No resource with id %d", msg.Id))
	}
	if msg.Creator != val.Creator {
		// unauthorized update
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Unauthorized update. Owner doesn't match")
	}

	k.SetResource(ctx, resource)

	return &types.MsgUpdateResourceResponse{}, nil
}
