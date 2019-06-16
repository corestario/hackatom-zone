#!/usr/bin/env bash

echo "Clearing previous files..."
rm -rf ~/.nft*

echo "Building..."
make install

echo "Initialization..."
nftd init hackNFT --chain-id NFTChain

echo "Adding keys..."

echo "Adding genesis account..."
nftcli keys add validator1 --recover <<< "12345678
base figure planet hazard sail easily honey advance tuition grab across unveil random kiss fence connect disagree evil recall latin cause brisk soft lunch
"

nftcli keys add alice --recover <<< "12345678
actor barely wait patrol moral amateur hole clerk misery truly salad wonder artefact orchard grit check abandon drip avoid shaft dirt thought melody drip
"

nftd add-genesis-account $(nftcli keys show validator1 -a) 1000token,100000000stake
nftd add-genesis-account $(nftcli keys show alice -a) 1000token

echo "Configuring..."
nftcli config chain-id NFTChain
nftcli config output json
nftcli config indent true
nftcli config trust-node true

nftd gentx --name validator1 <<< "12345678"
nftd collect-gentxs
nftd validate-genesis

echo "Starting node..."
nftd start