package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go/rpc"
	jsoniter "github.com/json-iterator/go"
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

	jsonMap := make(map[string]string)
	for i := 0; i < 100; i++ {
		k := fmt.Sprint(i)
		v := fmt.Sprintf("%d @#@!$!@#$!@#$@!#$!@#$DFQERQWERWQERQWERWQER", i)
		jsonMap[k] = v
	}
	jMStart := time.Now()
	jsonByte, err := json.Marshal(jsonMap)
	if err != nil {
		panic(err)
	}
	fmt.Println("encoding/json Marshal :", time.Now().Sub(jMStart))

	jUStart := time.Now()
	err = json.Unmarshal(jsonByte, &jsonMap)
	if err != nil {
		panic(err)
	}
	fmt.Println("encoding/json Unmarshal :", time.Now().Sub(jUStart))

	var json2 = jsoniter.ConfigCompatibleWithStandardLibrary

	j2MStart := time.Now()
	jsonByte, err = json2.Marshal(jsonMap)
	if err != nil {
		panic(err)
	}
	fmt.Println("jsoniter Marshal :", time.Now().Sub(j2MStart))

	j2UStart := time.Now()
	err = json2.Unmarshal(jsonByte, &jsonMap)
	if err != nil {
		panic(err)
	}
	fmt.Println("jsoniter Unmarshal :", time.Now().Sub(j2UStart))
}
