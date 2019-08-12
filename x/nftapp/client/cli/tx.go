package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/dgamingfoundation/hackatom-zone/x/nftapp/types"
	"github.com/spf13/cobra"
)

func GetCmdCreateNFT(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "createNFT [name] [description] [image] [tokenURI]",
		Short: "Creates NFT",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgCreateNFT(cliCtx.GetFromAddress(), args[0], args[1], args[2], args[3])
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdTransferTokenToHub(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "transfer [tokenURI] [price]",
		Short: "Transfers token to hub's marketplace",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			price, err := sdk.ParseCoin(args[1])
			if err != nil {
				return fmt.Errorf("failed to parse price: %v", err)
			}

			tokenURI := args[0]
			msg := types.NewMsgTransferTokenToHub(cliCtx.GetFromAddress(), tokenURI, price)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
