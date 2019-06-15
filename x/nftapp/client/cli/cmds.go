package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	// Group nftapp queries under a subcommand
	nftappQueryCmd := &cobra.Command{
		Use:   "nftapp",
		Short: "Querying commands for the nftapp module",
	}

	nftappQueryCmd.AddCommand(client.GetCommands(
		GetCmdGetNFTData(cdc),
		GetCmdNFTList(cdc),
	)...)

	return nftappQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	nftappTxCmd := &cobra.Command{
		Use:   "nftapp",
		Short: "nftapp transactions subcommands",
	}

	nftappTxCmd.AddCommand(client.PostCommands(
		GetCmdCreateNFT(cdc),
		GetCmdTransferTokenToHub(cdc),
	)...)

	return nftappTxCmd
}
