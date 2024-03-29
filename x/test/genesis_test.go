package test_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "zyx/testutil/keeper"
	"zyx/testutil/nullify"
	"zyx/x/test"
	"zyx/x/test/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DepositList: []types.Deposit{
			{
				TxHash: "0",
			},
			{
				TxHash: "1",
			},
		},
		WithdrawList: []types.Withdraw{
			{
				TxHash: "0",
			},
			{
				TxHash: "1",
			},
		},
		CallerGroupList: []types.CallerGroup{
			{
				Name: "0",
			},
			{
				Name: "1",
			},
		},
		SignerGroupList: []types.SignerGroup{
			{
				Name: "0",
			},
			{
				Name: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TestKeeper(t)
	test.InitGenesis(ctx, *k, genesisState)
	got := test.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DepositList, got.DepositList)
	require.ElementsMatch(t, genesisState.WithdrawList, got.WithdrawList)
	require.ElementsMatch(t, genesisState.CallerGroupList, got.CallerGroupList)
	require.ElementsMatch(t, genesisState.SignerGroupList, got.SignerGroupList)
	// this line is used by starport scaffolding # genesis/test/assert
}
