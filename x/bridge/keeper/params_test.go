package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "zyx/testutil/keeper"
	"zyx/x/bridge/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BridgeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
