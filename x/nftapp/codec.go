package nftapp

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on wire codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateNFT{}, "nftapp/CreateNFT", nil)
	cdc.RegisterConcrete(MsgTransferTokenToHub{}, "nftapp/TransferTokenToHub", nil)
	cdc.RegisterConcrete(BaseNFT{}, "nftapp/BaseNFT", nil)
}

var ModuleCdc = codec.New()
