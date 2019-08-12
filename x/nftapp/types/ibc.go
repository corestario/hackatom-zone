package types

import (
	"encoding/json"
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	xnft "github.com/cosmos/cosmos-sdk/x/nft"
)

type SellTokenPacket struct {
	Token *xnft.BaseNFT `json:"token"`
	Price sdk.Coin      `json:"price"`
}

func NewSellTokenPacket(token *xnft.BaseNFT, price sdk.Coin) *SellTokenPacket {
	return &SellTokenPacket{Token: token, Price: price}
}

func (m *SellTokenPacket) Timeout() uint64 {
	return math.MaxUint64
}

func (m *SellTokenPacket) Route() string {
	return "sell_token_packet"
}

func (m *SellTokenPacket) Commit() []byte {
	data, err := json.Marshal(m)
	if err != nil {
		panic(fmt.Errorf("failed to marshal SellTokenPacket packet: %v", err))
	}

	return data
}
