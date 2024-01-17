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

func createNCallerGroup(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.CallerGroup {
	items := make([]types.CallerGroup, n)
	for i := range items {
		items[i].Name = strconv.Itoa(i)

		keeper.SetCallerGroup(ctx, items[i])
	}
	return items
}

func TestCallerGroupGet(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNCallerGroup(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCallerGroup(ctx,
			item.Name,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCallerGroupRemove(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNCallerGroup(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCallerGroup(ctx,
			item.Name,
		)
		_, found := keeper.GetCallerGroup(ctx,
			item.Name,
		)
		require.False(t, found)
	}
}

func TestCallerGroupGetAll(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNCallerGroup(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCallerGroup(ctx)),
	)
}
