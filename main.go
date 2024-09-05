package main

import (
	"log"
	"xdchelp/abi"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Set the network you are running on
	// Options: "amoy", "apothem"
	// Default: "apothem"
	// Note: This is only used to determine the RPC URL
	runningOn := "apothem"

	// Set the RPC URL
	// Options: "https://erpc.apothem.network", "https://rpc-amoy.polygon.technology"
	// Note: This is only used to connect to the Ethereum client
	var rpcUrl string
	if runningOn == "amoy" {
		log.Default().Println("Running on Amoy (Polygon) network")
		rpcUrl = "https://rpc-amoy.polygon.technology"
	} else {
		log.Default().Println("Running on Apothem (XDC) network")
		//rpcUrl = "https://erpc.xinfin.network"
		//rpcUrl = "https://rpc.ankr.com/xdc"
		rpcUrl="https://devnetstats.apothem.network/devnet" //Devnet
		//rpcUrl="https://erpc.apothem.network"
	}

	// Set the contract address
	// Options: "0x14C32c1a52e83A7cB2BA02BcE44B5Be999D146b7" (Apothem Proxy), "0xd75b21da7748152c3b8e3D67032C6fdf01092fC8" (Amoy Proxy)
	// Note: This is the address of the deployed contract
	var contractAddress string
	if runningOn == "amoy" {
		log.Default().Println("Using Amoy contract address")
		//contractAddress = "0xd75b21da7748152c3b8e3D67032C6fdf01092fC8"
		contractAddress="0xF95BaF27B50c5Fe5F8662A4CEc9e721A27d40b88"
	} else {
		log.Default().Println("Using Apothem contract address")
		//contractAddress = "0x3279dbefabf3c6ac29d7ff24a6c46645f3f4403c"
		contractAddress="0x1796a4caf25f1a80626d8a2d26595b19b11697c9" //BBB Dev
		//contractAddress="0xe99500ab4a413164da49af83b9824749059b46ce"
	}

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
