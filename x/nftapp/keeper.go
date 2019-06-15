package nftapp

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/pkg/errors"
)

// StoreKey to be used when creating the KVStore
const StoreKey = "nft"

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc      *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// CreateNFT creates NFT in KVStore
func (k Keeper) CreateNFT(ctx sdk.Context, nft BaseNFT) {
	store := ctx.KVStore(k.storeKey)

	if store.Has([]byte(nft.GetID())) {
		return
	}

	store.Set([]byte(nft.GetID()), k.cdc.MustMarshalBinaryBare(nft))
}

func (k Keeper) DeleteNFT(ctx sdk.Context, sender sdk.AccAddress, nftID string) error {
	nft := k.GetNFT(ctx, nftID)
	if len(nft.ID) == 0 {
		return nil
	}

	if !nft.Owner.Equals(sender) {
		return errors.New("not an owner")
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(nftID))

	return nil
}

// GetNFT Gets the entire NFT metadata struct by id
func (k Keeper) GetNFT(ctx sdk.Context, id string) BaseNFT {
	store := ctx.KVStore(k.storeKey)

	if !store.Has([]byte(id)) {
		return BaseNFT{}
	}
	encodedNFT := store.Get([]byte(id))
	var nft BaseNFT
	k.cdc.MustUnmarshalBinaryBare(encodedNFT, &nft)
	return nft
}

// GetNFTList gets list of NFT tokens by owner's address
func (k Keeper) GetNFTList(ctx sdk.Context, owner sdk.AccAddress) NFTs {
	var (
		decodedNFT BaseNFT
	)
	nfts := NewNFTs()
	iterator := k.GetNFTIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		encodedNFT := iterator.Value()
		k.cdc.MustUnmarshalBinaryBare(encodedNFT, &decodedNFT)
		if decodedNFT.GetOwner().Equals(owner) {
			nfts.Add(NewNFTs(decodedNFT))
		}
	}
	return nfts
}

func (k Keeper) GetNFTIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}
