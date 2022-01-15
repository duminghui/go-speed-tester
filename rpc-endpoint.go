package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gagliardetto/solana-go/rpc"
)

type RpcInfo struct {
	Url    string `json:"url,omitempty"`
	Weight int    `json:"weight,omitempty"`
	Batch  bool   `json:"batch,omitempty"`
	Name   string `json:"name,omitempty"`
}

func (r RpcInfo) String() string {
	return fmt.Sprintf("{Url:%s, Weight:%d, Batch:%t, Name:%s}", r.Url, r.Weight, r.Batch, r.Name)
}

type Web3Config struct {
	Rpcs     []*RpcInfo `json:"rpcs,omitempty"`
	Strategy string     `json:"strategy,omitempty"`
	Success  bool       `json:"success,omitempty"`
}

func (w Web3Config) String() string {
	return fmt.Sprintf("{Rpcs:%s, Strategy:%s, Success:%t}", w.Rpcs, w.Strategy, w.Success)
}

var web3Config = &Web3Config{
	Rpcs: []*RpcInfo{
		//{Url: "https://raydium.rpcpool.com", Weight: 50},
		{Url: rpc.MainNetBeta_RPC, Weight: 50},
		{Url: "https://raydium.genesysgo.net", Weight: 30},
		{Url: "https://mainnet.rpcpool.com", Weight: 10},
		{Url: "https://free.rpcpool.com", Weight: 10},
		{Url: "https://solana-api.projectserum.com", Weight: 10},
		{Url: "https://solana-api.tt-prod.net", Weight: 10},
	},
	Strategy: "speed",
}

type speedInfo struct {
	url      string
	duration time.Duration
}

func getFastestEndpoint(rpcs []*RpcInfo) string {
	wg := &sync.WaitGroup{}
	var speedInfos []speedInfo
	for i := 0; i < 2; i++ {
		for _, rc := range rpcs {
			wg.Add(1)
			go func(rc *RpcInfo) {
				defer wg.Done()
				start := time.Now()
				out, err := rpc.New(rc.Url).GetRecentBlockhash(context.Background(), rpc.CommitmentConfirmed)
				if err != nil {
					if strings.Index(err.Error(), context.Canceled.Error()) < 0 {
						fmt.Printf("%s test speed error: %s\n", rc.Url, err)
					} else {
						fmt.Printf("Speed test cancel %s \n", rc.Url)
					}
					return
				}
				end := time.Now()
				duration := end.Sub(start)
				fmt.Printf("%35s: slot:%d, hash:%s,%s, %s\n", rc.Url, out.Context.Slot, out.Value.Blockhash,
					time.Now().Format("2006-01-02 15:04:05.999"),
					duration)
				speedInfos = append(speedInfos, speedInfo{
					url:      rc.Url,
					duration: end.Sub(start),
				})
			}(rc)
		}
		time.Sleep(500 * time.Millisecond)
	}
	wg.Wait()
	return speedInfos[0].url
}
