package main

import (
	"log"
	"xdchelp/abi"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//abigen --abi=./abi/grupo-soma.json --pkg=abi --out=grupo-soma.go
//Command to generate functions for ABI

func main() {

	var rpcUrl string
	//rpcUrl = "https://erpc.xinfin.network" //Mainnet
	rpcUrl="https://devnetstats.apothem.network/devnet" //Devnet
	//rpcUrl="https://erpc.apothem.network" //Apothem
	

	var contractAddress string
	contractAddress = "xdc3279dbefabf3c6ac29d7ff24a6c46645f3f4403c" //Mainnet
	contractAddress="0x1796a4caf25f1a80626d8a2d26595b19b11697c9" //Devnet
	//contractAddress="0xeDc8AD44C75FB093d1c515f6cB4Fa4a5e1448e63" //Apothem

	log.Default().Println("Connecting to Ethereum client")
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	log.Default().Println("Instantiating contract")
	address := common.HexToAddress(contractAddress)
	contract, err := abi.NewAbi(address, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a smart contract: %v", err)
	}

	log.Default().Println("Getting token URI from contract")
	tokenUri, err := contract.Name(nil);
	if err != nil {
		log.Fatalf("Failed to get token URI: %v", err)
	}

	log.Default().Printf("Token URI: %s\n", tokenUri)
}
