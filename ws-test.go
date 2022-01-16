package main

import (
	"context"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func wsTest() {
	wssUrl := "wss://free.rpcpool.com"
	client, err := ws.Connect(context.Background(), wssUrl)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	program := solana.MustPublicKeyFromBase58("DvVijge9HpEpfNVk8cTzdq2GuuF1eUcpVPGLf871HXFz") // serum
	fmt.Printf("wssUrl:%s\n", wssUrl)
	{
		// Subscribe to log events that mention the provided pubkey:
		sub, err := client.LogsSubscribeMentions(
			program,
			rpc.CommitmentRecent,
		)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Start LogsSubscribeMentions:%s\n", program)
		defer sub.Unsubscribe()

		for {
			got, err := sub.Recv()
			if err != nil {
				panic(err)
			}
			fmt.Println(time.Now())
			spew.Dump(got)
		}
	}
	if false {
		// Subscribe to all log events:
		sub, err := client.LogsSubscribe(
			ws.LogsSubscribeFilterAll,
			rpc.CommitmentRecent,
		)
		if err != nil {
			panic(err)
		}
		defer sub.Unsubscribe()

		for {
			got, err := sub.Recv()
			if err != nil {
				panic(err)
			}
			spew.Dump(got)
		}
	}
}
