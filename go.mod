module github.com/dgamingfoundation/hackatom-zone

go 1.12

require (
	github.com/cosmos/cosmos-sdk v0.28.2-0.20190616123619-7efb69cb3708
	github.com/gorilla/mux v1.7.0
	github.com/mattn/go-runewidth v0.0.4 // indirect
	github.com/olekukonko/tablewriter v0.0.1 // indirect
	github.com/pkg/errors v0.8.1
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.0.3
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/tendermint v0.31.5
)

replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5
