package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/dgamingfoundation/nftapp/x/nftapp/types"
	"github.com/gorilla/mux"
)

const (
	restName = "nftapp"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/nftapp/nft", createNFTHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/nftapp/nft/transfer", transferNFTHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/nftapp/nft/{%s}", getNFTHandler(cliCtx)).Methods("GET")
	r.HandleFunc("/nftapp/nft/list/{%s}/", getNFTListHandler(cliCtx)).Methods("GET")
}

func getNFTHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramType := vars[restName]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/nftapp/getNFTData/%s", paramType), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func getNFTListHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramType := vars[restName]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/nftapp/getNFTList/%s", paramType), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

type CreateNFTReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Owner   string       `json:"owner"`
	// User data
	Name     string `json:"name"`
	Password string `json:"password"`
	// Token data
	TokenName   string `json:"token_name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	TokenURI    string `json:"token_uri"`
}

func createNFTHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateNFTReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgCreateNFT(addr, req.TokenName, req.Description, req.Image, req.TokenURI)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		broadcastTransaction(cliCtx, w, baseReq, msg, req.Name, req.Password)
	}
}

type TransferToHubReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Owner   string       `json:"owner"`
	// User data
	Name     string `json:"name"`
	Password string `json:"password"`
	// Token data
	TokenID string `json:"token_id"`
	Price   string `json:"price"`
}

func transferNFTHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req TransferToHubReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		price, err := sdk.ParseCoin(req.Price)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgTransferTokenToHub(addr, req.TokenID, price)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		broadcastTransaction(cliCtx, w, baseReq, msg, req.Name, req.Password)
	}
}

func broadcastTransaction(
	cliCtx context.CLIContext,
	w http.ResponseWriter,
	bq rest.BaseReq,
	msg sdk.Msg,
	name,
	password string) {
	gasAdj, ok := rest.ParseFloat64OrReturnBadRequest(w, bq.GasAdjustment, flags.DefaultGasAdjustment)
	if !ok {
		return
	}

	_, gas, err := flags.ParseGas(bq.Gas)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	txBldr := authtxb.NewTxBuilder(
		utils.GetTxEncoder(cliCtx.Codec), bq.AccountNumber, bq.Sequence, gas, gasAdj,
		bq.Simulate, bq.ChainID, bq.Memo, bq.Fees, bq.GasPrices,
	)

	msgBytes, err := txBldr.BuildAndSign(name, password, []sdk.Msg{msg})
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := cliCtx.BroadcastTxCommit(msgBytes)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	rest.PostProcessResponse(w, cliCtx, resp)
}
