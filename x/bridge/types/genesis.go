package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		CallerList: []Caller{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in caller
	callerIdMap := make(map[uint64]bool)
	callerCount := gs.GetCallerCount()
	for _, elem := range gs.CallerList {
		if _, ok := callerIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for caller")
		}
		if elem.Id >= callerCount {
			return fmt.Errorf("caller id should be lower or equal than the last id")
		}
		callerIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
