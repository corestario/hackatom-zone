package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/dgamingfoundation/hackatom-zone/x/nftapp/types"
	"github.com/spf13/cobra"
)

func GetCmdGetNFTData(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "getNFTData [id]",
		Short: "Get NFT data by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			id := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/nftapp/getNFTData/%s", id), nil)
			if err != nil {
				fmt.Printf("could not get NFT data  - %s \n", err.Error())
				return nil
			}

			var out types.BaseNFT
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdNFTList(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "getNFTList [address]",
		Short: "Get NFT list by owner's address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			address := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/nftapp/getNFTList/%s", address), nil)
			if err != nil {
				fmt.Printf("could not get NFT list - %s \n", err.Error())
				return nil
			}

			var out types.NFTs
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
