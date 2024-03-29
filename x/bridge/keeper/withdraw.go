package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"zyx/x/bridge/types"
)

// SetWithdraw set a specific withdraw in the store from its index
func (k Keeper) SetWithdraw(ctx sdk.Context, withdraw types.Withdraw) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawKeyPrefix))
	b := k.cdc.MustMarshal(&withdraw)
	store.Set(types.WithdrawKey(
		withdraw.Index,
	), b)
}

// GetWithdraw returns a withdraw from its index
func (k Keeper) GetWithdraw(
	ctx sdk.Context,
	index string,

) (val types.Withdraw, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawKeyPrefix))

	b := store.Get(types.WithdrawKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveWithdraw removes a withdraw from the store
func (k Keeper) RemoveWithdraw(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawKeyPrefix))
	store.Delete(types.WithdrawKey(
		index,
	))
}

// GetAllWithdraw returns all withdraw
func (k Keeper) GetAllWithdraw(ctx sdk.Context) (list []types.Withdraw) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Withdraw
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
