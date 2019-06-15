package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/dgamingfoundation/nftapp/x/nftapp"
	"github.com/spf13/cobra"
)

func GetCmdGetNFTData(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "getNFTData [id]",
		Short: "Get NFT data by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			id := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/getNFTData/%s", queryRoute, id), nil)
			if err != nil {
				fmt.Printf("could not get NFT data  - %s \n", err.Error())
				return nil
			}

			var out nftapp.BaseNFT
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdNFTList(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "getNFTList [address]",
		Short: "Get NFT list by owner's address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			address := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/getNFTList/%s", queryRoute, address), nil)
			if err != nil {
				fmt.Printf("could not get NFT list - %s \n", err.Error())
				return nil
			}

			var out nftapp.NFTs
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}