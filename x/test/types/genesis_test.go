package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"zyx/x/test/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated deposit",
			genState: &types.GenesisState{
				DepositList: []types.Deposit{
					{
						TxHash: "0",
					},
					{
						TxHash: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated withdraw",
			genState: &types.GenesisState{
				WithdrawList: []types.Withdraw{
					{
						TxHash: "0",
					},
					{
						TxHash: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
