module github.com/dgamingfoundation/hackatom-zone

go 1.12

require (
	github.com/cosmos/cosmos-sdk v0.28.2-0.20190616123619-7efb69cb3708
	github.com/gorilla/mux v1.7.0
	github.com/pkg/errors v0.8.1
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.3.2
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/tendermint v0.32.2
	github.com/tendermint/tm-db v0.1.1
)

replace github.com/cosmos/cosmos-sdk => github.com/dgamingfoundation/cosmos-sdk v0.0.0-20190806155809-7f4388fe7599
