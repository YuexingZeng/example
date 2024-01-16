package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "zyx/testutil/keeper"
	"zyx/x/test/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TestKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
