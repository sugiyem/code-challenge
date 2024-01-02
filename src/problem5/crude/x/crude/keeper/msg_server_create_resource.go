package keeper

import (
	"context"

	"crude/x/crude/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateResource(goCtx context.Context, msg *types.MsgCreateResource) (*types.MsgCreateResourceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var resource = types.Resource{
		Metadata: msg.Metadata,
		Value:    msg.Value,
		Creator:  msg.Creator,
	}
	id := k.AppendResource(ctx, resource)

	return &types.MsgCreateResourceResponse{Id: id}, nil
}
