package keeper

import (
	"context"
	"strings"

	"crude/x/crude/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListResource(goCtx context.Context, req *types.QueryListResourceRequest) (*types.QueryListResourceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ResourceKey))

	var resources []types.Resource

	pageRes, err := query.FilteredPaginate(store, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var resource types.Resource

		if err := k.cdc.Unmarshal(value, &resource); err != nil {
			return false, err
		}

		var cond = true
		if !strings.Contains(resource.Metadata, req.MetadataFilter) {
			cond = false
		}
		if resource.Value < req.ValueLow || resource.Value > req.ValueHigh {
			cond = false
		}

		if !cond {
			return false, nil
		}

		resources = append(resources, resource)
		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListResourceResponse{Resource: resources, Pagination: pageRes}, nil
}
