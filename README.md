## Install

Run 

```bash
./run.sh
```

and follow the prompts. See `./run.sh` for details on how to create users, etc.

# REST API

Start the rest server with:

```bash
nftcli rest-server --chain-id NFTChain --trust-node
```

#### Create Token

```
curl -s -XPOST http://localhost:1317/nftapp/nft --data-binary '{"base_req":{"from":"cosmos1mfgzjrd5klx3saahfy0rgf7ec9utdjx56y6smy","chain_id":"NFTChain", "account_number": "0", "sequence": "1" },"token_name":"alpha","description":"beta","image":"gamma","token_uri":"kappa","owner":"cosmos1mfgzjrd5klx3saahfy0rgf7ec9utdjx56y6smy", "name": "validator1", "password": "Committed"}'
```

*Output:*

```
{
  "height": "207",
  "txhash": "B237D9692E0CC00C0A1FB96DF41A19BF12F7DC4F78DE27EFF8280D34AF1A7DCE",
  "raw_log": "[{\"msg_index\":0,\"success\":true,\"log\":\"\"}]",
  "logs": [
    {
      "msg_index": 0,
      "success": true,
      "log": ""
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "14219",
  "tags": [
    {
      "key": "action",
      "value": "create_nft"
    }
  ]
}
```

#### Show tokens for some address

```
curl -s http://localhost:1317/nftapp/nft/list/cosmos1np0wt2u949r3k6r5km023c6w5vp5dedyrglld9/
```

*Output*:

```
{
  "81bc9d76-5a1d-44d3-bdd5-c8a4bba155b1": {
    "id": "81bc9d76-5a1d-44d3-bdd5-c8a4bba155b1",
    "owner": "cosmos1np0wt2u949r3k6r5km023c6w5vp5dedyrglld9",
    "name": "\"alpha\"",
    "description": "\"beta\"",
    "image": "\"gamma\"",
    "token_uri": "\"kappa\""
  }
}
```

#### Show Token Info

```
curl -s http://localhost:1317/nftapp/nft/81bc9d76-5a1d-44d3-bdd5-c8a4bba155b1
```

*Output:*

```
{
  "type": "nftapp/BaseNFT",
  "value": {
    "id": "81bc9d76-5a1d-44d3-bdd5-c8a4bba155b1",
    "owner": "cosmos1np0wt2u949r3k6r5km023c6w5vp5dedyrglld9",
    "name": "alpha",
    "description": "beta",
    "image": "gamma",
    "token_uri": "kappa"
  }
}
```

#### Transfer Token

```
curl -s -XPOST http://localhost:1317/nftapp/nft/transfer --data-binary '{"base_req":{"from":"cosmos1mfgzjrd5klx3saahfy0rgf7ec9utdjx56y6smy","chain_id":"NFTChain", "account_number": "0", "sequence": "4" },"token_id":"81bc9d76-5a1d-44d3-bdd5-c8a4bba155b1","price":"100","owner":"cosmos1mfgzjrd5klx3saahfy0rgf7ec9utdjx56y6smy", "name": "validator1", "password": "Committed","price":"10token"}'
```

*Output:*

```
{
  "height": "219",
  "txhash": "263CC0B2B49A382168EB1B5E5888630FD89FEE577D07FFD9C49076B7B2486786",
  "raw_log": "[{\"msg_index\":0,\"success\":true,\"log\":\"\"}]",
  "logs": [
    {
      "msg_index": 0,
      "success": true,
      "log": ""
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "11872",
  "tags": [
    {
      "key": "action",
      "value": "transfer_token_to_hub"
    }
  ]
}
```

# CLI

* Create NFT with a name, description, image and token_uri for Jack's account
```bash
nftcli tx nftapp createNFT NAME DESCRIPTION IMAGE TOKEN_URI --from validator1
```
##### Output
```bash
{
  "chain_id": "lol",
  "account_number": "0",
  "sequence": "3",
  "fee": {
    "amount": [],
    "gas": "200000"
  },
  "msgs": [
    {
      "type": "nftapp/CreateNFT",
      "value": {
        "Owner": "cosmos1xm697yarpz6lhh265ngaxwgrvygpkpkxvce7vs",
        "TokenURI": "TOKEN_URI",
        "Description": "DESCRIPTION",
        "Image": "IMAGE",
        "Name": "NAME"
      }
    }
  ],
  "memo": ""
}

confirm transaction before signing and broadcasting [Y/n]: y
Password to sign with 'jack':
{
  "height": "0",
  "txhash": "14B72D74A4670B1B880B835063F3C9B3844773C48214D6F733A1026E8A2A70A5"
}
```

* Get a list of NFTs for Jack's account

```bash
nftcli query nftapp getNFTList $(nftcli keys show validator1 -a)
```
##### Output
```json
{
    "7724138b-40ab-4401-b4e8-3afcbd9adb8b": {
        "id": "7724138b-40ab-4401-b4e8-3afcbd9adb8b",
        "owner": "cosmos1xm697yarpz6lhh265ngaxwgrvygpkpkxvce7vs",
        "name": "\"TOKEN_URI\"",
        "description": "\"DESCRIPTION\"",
        "image": "\"IMAGE\"",
        "token_uri": "\"NAME\""
      }
}
```

* Get NFT data by ID
```bash
nftcli query nftapp getNFTData 443a901b-5738-475c-91a2-11cc24af9e01
```

##### Output
```json
{
  "type": "nftapp/BaseNFT",
  "value": {
    "id": "7724138b-40ab-4401-b4e8-3afcbd9adb8b",
    "owner": "cosmos1xm697yarpz6lhh265ngaxwgrvygpkpkxvce7vs",
    "name": "TOKEN_URI",
    "description": "DESCRIPTION",
    "image": "IMAGE",
    "token_uri": "NAME"
  }
}

```


nftcli tx nftapp transfer 7cd2ea3b-1283-4b30-b028-c840803e39d7 10token --from validator1
