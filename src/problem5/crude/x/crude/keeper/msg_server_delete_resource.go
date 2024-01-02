package keeper

import (
	"context"

	"crude/x/crude/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteResource(goCtx context.Context, msg *types.MsgDeleteResource) (*types.MsgDeleteResourceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeleteResourceResponse{}, nil
}
