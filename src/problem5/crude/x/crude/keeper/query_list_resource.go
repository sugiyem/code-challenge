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
	var count uint64 = 0

	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var resource types.Resource
		if err := k.cdc.Unmarshal(value, &resource); err != nil {
			return err
		}

		var cond = true
		if !strings.Contains(resource.Metadata, req.MetadataFilter) {
			cond = false
		}
		if resource.Value < req.ValueLow || resource.Value > req.ValueHigh {
			cond = false
		}

		if cond {
			resources = append(resources, resource)
			count++
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pageRes.Total = count

	return &types.QueryListResourceResponse{Resource: resources, Pagination: pageRes}, nil
}
