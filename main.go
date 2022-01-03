package main

import (
	"context"
	"fmt"
	"github.com/gagliardetto/solana-go/rpc"
	"time"
)

func main() {
	fmt.Println("Test 1.........")
	endpoint := getFastestEndpoint(web3Config.Rpcs)
	time.Sleep(3 * time.Second)
	fmt.Println("Test 2......")
	getFastestEndpoint2(web3Config.Rpcs)
	time.Sleep(3 * time.Second)

	rcp := rpc.New(endpoint)
	start := time.Now()
	blockhash, err := rcp.GetRecentBlockhash(context.Background(), rpc.CommitmentConfirmed)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("blockhash %s, Use time:%s\n", blockhash.Value.Blockhash, time.Now().Sub(start))
	fmt.Println("End....")
}
