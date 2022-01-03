package main

import (
	"context"
	"fmt"
	rpc2 "github.com/portto/solana-go-sdk/rpc"
	"sort"
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
		{Url: "https://solana-api.projectserum.com", Weight: 10},
		{Url: "https://solana-api.tt-prod.net", Weight: 10},
	},
	Strategy: "speed",
}

func getWeightEndpoint(rpcs []*RpcInfo) string {
	sort.Slice(rpcs, func(i, j int) bool {
		return rpcs[i].Weight > rpcs[j].Weight
	})
	return rpcs[0].Url
}

type speedInfo struct {
	url      string
	duration time.Duration
}

func getFastestEndpoint(rpcs []*RpcInfo) string {
	wg := &sync.WaitGroup{}
	var speedInfos []speedInfo
	wg.Add(1)
	for _, rc := range rpcs {
		go func(rc *RpcInfo) {
			start := time.Now()
			out, err := rpc.New(rc.Url).GetEpochInfo(context.Background(), rpc.CommitmentConfirmed)
			if err != nil {
				fmt.Printf("%s test speed error: %s\n", rc.Url, err)
				return
			}
			end := time.Now()
			duration := end.Sub(start)
			fmt.Printf("%s: %d %s\n", rc.Url, out.BlockHeight, duration)
			speedInfos = append(speedInfos, speedInfo{
				url:      rc.Url,
				duration: end.Sub(start),
			})
			if len(speedInfos) == 1 {
				wg.Done()
			}
		}(rc)
	}
	wg.Wait()
	return speedInfos[0].url
}

func getFastestEndpoint2(rpcs []*RpcInfo) string {
	wg := &sync.WaitGroup{}
	var speedInfos []speedInfo
	wg.Add(1)
	for _, rc := range rpcs {
		go func(rc *RpcInfo) {
			start := time.Now()
			rpcClient := rpc2.NewRpcClient(rc.Url)
			out, err := rpcClient.GetEpochInfoWithConfig(context.Background(), rpc2.GetEpochInfoConfig{
				Commitment: rpc2.CommitmentConfirmed,
			})
			if err != nil {
				fmt.Printf("%s test speed error: %s\n", rc.Url, err)
				return
			}
			end := time.Now()
			duration := end.Sub(start)
			fmt.Printf("2 %s: %d, %d, %s\n", rc.Url, out.ID, out.Result.Epoch, duration)
			speedInfos = append(speedInfos, speedInfo{
				url:      rc.Url,
				duration: end.Sub(start),
			})
			if len(speedInfos) == 1 {
				wg.Done()
			}
		}(rc)
	}
	wg.Wait()
	return speedInfos[0].url
}

func getRpcEndpoint() string {
	useWeb3Config := web3Config
	if useWeb3Config.Strategy == "weight" {
		return getWeightEndpoint(useWeb3Config.Rpcs)
	} else {
		return getFastestEndpoint(useWeb3Config.Rpcs)
	}
}
