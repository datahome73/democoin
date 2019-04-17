package democoin

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on wire codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetName{}, "democoin/SetName", nil)
	cdc.RegisterConcrete(MsgBuyName{}, "democoin/BuyName", nil)
}