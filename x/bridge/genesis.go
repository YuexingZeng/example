package bridge

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"zyx/x/bridge/keeper"
	"zyx/x/bridge/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the caller
	for _, elem := range genState.CallerList {
		k.SetCaller(ctx, elem)
	}

	// Set caller count
	k.SetCallerCount(ctx, genState.CallerCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.CallerList = k.GetAllCaller(ctx)
	genesis.CallerCount = k.GetCallerCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
