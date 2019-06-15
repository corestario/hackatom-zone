package nftapp

import (
	"github.com/cosmos/cosmos-sdk/client"
	nftappcmd "github.com/dgamingfoundation/nftapp/x/nftapp/client/cli"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	// Group nftapp queries under a subcommand
	nftappQueryCmd := &cobra.Command{
		Use:   "nftapp",
		Short: "Querying commands for the nftapp module",
	}

	nftappQueryCmd.AddCommand(client.GetCommands(
		nftappcmd.GetCmdGetNFTData(mc.storeKey, mc.cdc),
		nftappcmd.GetCmdNFTList(mc.storeKey, mc.cdc),
	)...)

	return nftappQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	nftappTxCmd := &cobra.Command{
		Use:   "nftapp",
		Short: "nftapp transactions subcommands",
	}

	nftappTxCmd.AddCommand(client.PostCommands(
		nftappcmd.GetCmdCreateNFT(mc.cdc),
		nftappcmd.GetCmdTransferTokenToHub(mc.cdc),
	)...)

	return nftappTxCmd
}
