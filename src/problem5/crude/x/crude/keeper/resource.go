package keeper

import (
	"encoding/binary"

	"crude/x/crude/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendResource(ctx sdk.Context, resource types.Resource) uint64 {
	count := k.GetResourceCount(ctx)
	resource.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ResourceKey))
	appendedValue := k.cdc.MustMarshal(&resource)

	store.Set(GetResourceIDBytes(resource.Id), appendedValue)
	k.SetResourceCount(ctx, count+1)

	return count
}

func (k Keeper) GetResource(ctx sdk.Context, id uint64) (val types.Resource, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ResourceKey))
	b := store.Get(GetResourceIDBytes(id))

	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

func (k Keeper) GetResourceCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ResourceCountKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetResource(ctx sdk.Context, post types.Resource) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ResourceKey))
	b := k.cdc.MustMarshal(&post)

	store.Set(GetResourceIDBytes(post.Id), b)
}

func (k Keeper) SetResourceCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ResourceCountKey)
	bz := make([]byte, 8)

	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func GetResourceIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) RemoveResource(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ResourceKey))
	store.Delete(GetResourceIDBytes(id))
}
