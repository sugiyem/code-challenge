package keeper

import (
	"context"
	"fmt"

	"crude/x/crude/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteResource(goCtx context.Context, msg *types.MsgDeleteResource) (*types.MsgDeleteResourceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetResource(ctx, msg.Id)
	if !found {
		// there is no resource with given id
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("No resource with id %d", msg.Id))
	}
	if msg.Creator != val.Creator {
		// unauthorized delete
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Unauthorized delete. Owner doesn't match")
	}

	k.RemoveResource(ctx, msg.Id)

	return &types.MsgDeleteResourceResponse{}, nil
}
