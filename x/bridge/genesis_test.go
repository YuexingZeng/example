package bridge_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "zyx/testutil/keeper"
	"zyx/testutil/nullify"
	"zyx/x/bridge"
	"zyx/x/bridge/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CallerList: []types.Caller{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		CallerCount: 2,
		DepositList: []types.Deposit{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		WithdrawList: []types.Withdraw{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BridgeKeeper(t)
	bridge.InitGenesis(ctx, *k, genesisState)
	got := bridge.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CallerList, got.CallerList)
	require.Equal(t, genesisState.CallerCount, got.CallerCount)
	require.ElementsMatch(t, genesisState.DepositList, got.DepositList)
	require.ElementsMatch(t, genesisState.WithdrawList, got.WithdrawList)
	// this line is used by starport scaffolding # genesis/test/assert
}
