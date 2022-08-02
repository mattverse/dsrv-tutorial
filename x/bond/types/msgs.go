package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgSendMoney = "send_money"
)

var (
	_ sdk.Msg = &MsgSendMoney{}
)

func NewMsgSendMoney(bonder sdk.AccAddress, coin sdk.Coin) *MsgSendMoney {
	return &MsgSendMoney{
		Sender: bonder.String(),
		Coin:   coin,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgSendMoney) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgSendMoney) Type() string { return TypeMsgSendMoney }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgSendMoney) GetSigners() []sdk.AccAddress {
	bonder, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{bonder}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgSendMoney) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgSendMoney) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}

	return msg.Coin.Validate()
}
