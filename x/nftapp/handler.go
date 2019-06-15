package nftapp

import (
	"fmt"

	"github.com/dgamingfoundation/nftapp/x/nftapp/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	uuid "github.com/satori/go.uuid"
)

// NewHandler returns a handler for "nft" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgCreateNFT:
			return handleMsgCreateNFT(ctx, keeper, msg)
		case types.MsgTransferTokenToHub:
			return handleMsgTransferTokenToHub(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nftapp Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to create nft
func handleMsgCreateNFT(ctx sdk.Context, keeper Keeper, msg types.MsgCreateNFT) sdk.Result {
	id := uuid.NewV4()
	nft := types.NewBaseNFT(id.String(), msg.Owner, msg.Name, msg.Description, msg.Image, msg.TokenURI)
	keeper.CreateNFT(ctx, nft)
	return sdk.Result{}
}

// Handle a message to create nft
func handleMsgTransferTokenToHub(ctx sdk.Context, keeper Keeper, msg types.MsgTransferTokenToHub) sdk.Result {
	// TODO: implement IBC transfer.

	if err := keeper.DeleteNFT(ctx, msg.Owner, msg.TokenURI); err != nil {
		return sdk.ErrUnauthorized(fmt.Sprintf("failed to delete token: %v", err.Error())).Result()
	}

	return sdk.Result{}
}
