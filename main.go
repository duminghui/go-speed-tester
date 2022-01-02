package main

import (
	"context"
	"fmt"
	"github.com/gagliardetto/solana-go/rpc"
	"time"
)

func main() {
	endpoint := getRpcEndpoint()
	rcp := rpc.New(endpoint)
	start := time.Now()
	blockhash, err := rcp.GetRecentBlockhash(context.Background(), rpc.CommitmentConfirmed)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("blockhash %s, Use time:%s\n", blockhash.Value.Blockhash, time.Now().Sub(start))
	time.Sleep(5 * time.Second)
	fmt.Println("End....")
}
