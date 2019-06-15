package nftapp

import (
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the nftapp Querier
const (
	QueryGetNFTData = "getNFTData"
	QueryGetNFTList = "getNFTList"

)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryGetNFTData:
			return queryGetNFTData(ctx, path[1:], req, keeper)
		case QueryGetNFTList:
			return queryGetNFTList(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown nftapp query endpoint")
		}
	}
}

// nolint: unparam
func queryGetNFTData(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	id := path[0]

	nft := keeper.GetNFT(ctx, id)

	if nft.GetID() == "" {
		return []byte{}, sdk.ErrUnknownRequest("could not find NFT")
	}

	bz, err := codec.MarshalJSONIndent(keeper.cdc, nft)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// nolint: unparam
func queryGetNFTList(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	address, err := sdk.AccAddressFromBech32(path[0])

	if err != nil {
		return nil, sdk.ErrInvalidAddress(err.Error())
	}

	nfts := keeper.GetNFTList(ctx, address)

	bz, err := codec.MarshalJSONIndent(keeper.cdc, nfts)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}