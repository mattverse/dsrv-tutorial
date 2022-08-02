package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mattverse/dsrv-tutorial/x/bond/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the bond MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) SendMoney(goCtx context.Context, msg *types.MsgSendMoney) (*types.MsgSendMoneyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	bonder, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = k.keeper.TakeMoneyFromUser(ctx, bonder, msg.Coin)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendMoneyResponse{}, nil
}
