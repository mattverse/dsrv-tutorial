package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sapiens-cosmos/arbiter/x/bond/types"
)

// setParams sets the total set of bond parameters.
func (k Keeper) setParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}