package types

// NewGenesisState creates a new GenesisState instanc
func NewGenesisState(baseDenom string) *GenesisState {
	return &GenesisState{}
}

// DefaultGenesisState gets the raw genesis raw message for testing
func DefaultGenesisState() *GenesisState {
	return NewGenesisState("stake")
}

func (genesis GenesisState) Validate() error {
	// TODO
	return nil
}
