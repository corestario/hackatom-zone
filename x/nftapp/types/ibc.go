package types

import (
	"encoding/json"
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SellTokenPacket struct {
	Token *BaseNFT `json:"token"`
	Price sdk.Coin `json:"price"`
}

func NewSellTokenPacket(token *BaseNFT, price sdk.Coin) *SellTokenPacket {
	return &SellTokenPacket{Token: token, Price: price}
}

func (m *SellTokenPacket) Timeout() uint64 {
	return math.MaxUint64
}

func (m *SellTokenPacket) Commit() []byte {
	data, err := json.Marshal(m)
	if err != nil {
		panic(fmt.Errorf("failed to marshal SellTokenPacket packet: %v", err))
	}

	return data
}
