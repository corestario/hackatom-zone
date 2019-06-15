package nftapp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dgamingfoundation/nftapp/x/nftapp/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewGenesisState(nfts []types.NFT) types.GenesisState {
	return types.GenesisState{NFTS: nil}
}

func ValidateGenesis(data types.GenesisState) error {
	for range data.NFTS {
		// TODO: validates
	}
	return nil
}

func DefaultGenesisState() types.GenesisState {
	return types.GenesisState{
		NFTS: types.NewNFTs(),
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data types.GenesisState) []abci.ValidatorUpdate {
	for _, nft := range data.NFTS {
		keeper.CreateNFT(ctx, nft)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) types.GenesisState {
	var nfts types.NFTs
	iterator := k.GetNFTIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		id := string(iterator.Key())
		var nft *types.BaseNFT
		nft, err := k.GetNFT(ctx, id)
		if err != nil {
			continue
		}
		nfts.Add(types.NewNFTs(*nft))
	}
	return types.GenesisState{NFTS: nfts}
}
