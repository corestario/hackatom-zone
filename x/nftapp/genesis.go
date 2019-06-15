package nftapp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewGenesisState(nfts []NFT) GenesisState {
	return GenesisState{NFTS: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, _ = range data.NFTS {
		// TODO: validates
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		NFTS: NewNFTs(),
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, nft := range data.NFTS {
		keeper.CreateNFT(ctx, nft)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var nfts NFTs
	iterator := k.GetNFTIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		id := string(iterator.Key())
		var nft BaseNFT
		nft = k.GetNFT(ctx, id)
		nfts.Add(NewNFTs(nft))
	}
	return GenesisState{NFTS: nfts}
}
