package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/mattverse/dsrv-tutorial/x/bond/types"
)

// keeper of the bond store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryMarshaler
	bankKeeper types.BankKeeper
	paramstore paramtypes.Subspace

	treasuryModuleName string
	treasuryKeeper     types.TreasuryKeeper
}

// NewKeeper creates a new bond Keeper instance
func NewKeeper(cdc codec.BinaryMarshaler, key sdk.StoreKey, bk types.BankKeeper, ps paramtypes.Subspace) Keeper {
	return Keeper{
		storeKey:   key,
		cdc:        cdc,
		bankKeeper: bk,
		paramstore: ps,
	}
}
