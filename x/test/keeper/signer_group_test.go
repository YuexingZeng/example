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

func createNSignerGroup(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SignerGroup {
	items := make([]types.SignerGroup, n)
	for i := range items {
		items[i].Name = strconv.Itoa(i)

		keeper.SetSignerGroup(ctx, items[i])
	}
	return items
}

func TestSignerGroupGet(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNSignerGroup(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSignerGroup(ctx,
			item.Name,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSignerGroupRemove(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNSignerGroup(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSignerGroup(ctx,
			item.Name,
		)
		_, found := keeper.GetSignerGroup(ctx,
			item.Name,
		)
		require.False(t, found)
	}
}

func TestSignerGroupGetAll(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNSignerGroup(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSignerGroup(ctx)),
	)
}
