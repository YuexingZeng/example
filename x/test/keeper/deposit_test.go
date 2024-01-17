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

func createNDeposit(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Deposit {
	items := make([]types.Deposit, n)
	for i := range items {
		items[i].TxHash = strconv.Itoa(i)

		keeper.SetDeposit(ctx, items[i])
	}
	return items
}

func TestDepositGet(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNDeposit(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDeposit(ctx,
			item.TxHash,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDepositRemove(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNDeposit(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDeposit(ctx,
			item.TxHash,
		)
		_, found := keeper.GetDeposit(ctx,
			item.TxHash,
		)
		require.False(t, found)
	}
}

func TestDepositGetAll(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNDeposit(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDeposit(ctx)),
	)
}
