package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "zyx/testutil/keeper"
	"zyx/testutil/nullify"
	"zyx/x/test/keeper"
	"zyx/x/test/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNWithdraw(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Withdraw {
	items := make([]types.Withdraw, n)
	for i := range items {
		items[i].TxHash = strconv.Itoa(i)

		keeper.SetWithdraw(ctx, items[i])
	}
	return items
}

func TestWithdrawGet(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNWithdraw(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetWithdraw(ctx,
			item.TxHash,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestWithdrawRemove(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNWithdraw(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWithdraw(ctx,
			item.TxHash,
		)
		_, found := keeper.GetWithdraw(ctx,
			item.TxHash,
		)
		require.False(t, found)
	}
}

func TestWithdrawGetAll(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNWithdraw(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllWithdraw(ctx)),
	)
}
