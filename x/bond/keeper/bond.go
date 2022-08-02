package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mattverse/dsrv-tutorial/x/bond/types"
)

func (k Keeper) TakeMoneyFromUser(ctx sdk.Context, bonder sdk.AccAddress, coin sdk.Coin) error {

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, bonder, types.ModuleName, sdk.NewCoins(coin))
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) GetModuleBalance(ctx sdk.Context) (sdk.Coin, error) {
	supply := k.bankKeeper.GetSupply(ctx)
	return supply.GetTotal()[0], nil
}
