#!/usr/bin/env bash

echo "Clearing previous files..."
rm -rf ~/.nft*

echo "Building..."
make install

echo "Initialization..."
nftd init hackNFT --chain-id NFTChain

echo "Adding keys..."
nftcli keys add validator1

echo "Adding genesis account..."
nftd add-genesis-account $(nftcli keys show validator1 -a) 1000token,100000000stake

echo "Configuring..."
nftcli config chain-id NFTChain
nftcli config output json
nftcli config indent true
nftcli config trust-node true
nftcli config connection-id market_connection
nftcli config counterparty-id hub
nftcli config counterparty-client-id me

nftd gentx --name validator1
nftd collect-gentxs
nftd validate-genesis

echo "Starting node..."
nftd start