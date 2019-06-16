package nftapp

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/dgamingfoundation/hackatom-zone/x/nftapp/types"
)

// RegisterCodec registers concrete types on wire codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(types.MsgCreateNFT{}, "nftapp/CreateNFT", nil)
	cdc.RegisterConcrete(types.MsgTransferTokenToHub{}, "nftapp/TransferTokenToHub", nil)
	cdc.RegisterConcrete(types.BaseNFT{}, "nftapp/BaseNFT", nil)
}

var ModuleCdc = codec.New()
