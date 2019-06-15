package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const ModuleName = "nftapp"
const RouterKey = ModuleName

// --------------------------------------------------------------------------
//
// CreateNFT
//
// --------------------------------------------------------------------------

// MsgCreateNFT defines a CreateNFT message
type MsgCreateNFT struct {
	Owner       sdk.AccAddress
	TokenURI    string
	Description string
	Image       string
	Name        string
}

// NewMsgCreateNFT is a constructor function for MsgCreateNFT
func NewMsgCreateNFT(owner sdk.AccAddress, name, description, image, tokenURI string) MsgCreateNFT {
	return MsgCreateNFT{
		Owner:       owner,
		TokenURI:    tokenURI,
		Description: description,
		Image:       image,
		Name:        name,
	}
}

// Route should return the name of the module
func (msg MsgCreateNFT) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateNFT) Type() string { return "create_nft" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateNFT) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateNFT) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgCreateNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// --------------------------------------------------------------------------
//
// TransferTokenToHub
//
// --------------------------------------------------------------------------

// MsgTransferTokenToHub defines a TransferTokenToHub message
type MsgTransferTokenToHub struct {
	Owner   sdk.AccAddress
	TokenID string
	Price   sdk.Coin
}

// NewMsgCreateNFT is a constructor function for MsgCreateNFT
func NewMsgTransferTokenToHub(owner sdk.AccAddress, tokenURI string, price sdk.Coin) MsgTransferTokenToHub {
	return MsgTransferTokenToHub{
		Owner:   owner,
		TokenID: tokenURI,
		Price:   price,
	}
}

// Route should return the name of the module
func (msg MsgTransferTokenToHub) Route() string { return RouterKey }

// Type should return the action
func (msg MsgTransferTokenToHub) Type() string { return "transfer_token_to_hub" }

// ValidateBasic runs stateless checks on the message
func (msg MsgTransferTokenToHub) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgTransferTokenToHub) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgTransferTokenToHub) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
