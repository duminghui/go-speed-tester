package rpchelp

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func GetMultipleAccountsWithOpts(rpcClient *rpc.Client, ctx context.Context, accounts []solana.PublicKey,
	opts *rpc.GetMultipleAccountsOpts) ([]*rpc.KeyedAccount, error) {
	// rpcClient.GetMultipleAccountsWithOpts max 100 account
	accountLen := len(accounts)
	size := accountLen / 100
	if accountLen%100 != 0 {
		size++
	}
	accountsSplit := make([][]solana.PublicKey, 0, size)
	accountsSub := make([]solana.PublicKey, 0, 100)
	for _, account := range accounts {
		if len(accountsSub) >= 100 {
			accountsSplit = append(accountsSplit, accountsSub)
			accountsSub = make([]solana.PublicKey, 0, 100)
		}
		accountsSub = append(accountsSub, account)
	}
	if len(accountsSub) > 0 {
		accountsSplit = append(accountsSplit, accountsSub)
	}
	out := make([]*rpc.KeyedAccount, 0, accountLen)
	for index1, accs := range accountsSplit {
		result, err := rpcClient.GetMultipleAccountsWithOpts(ctx, accs, opts)
		if err != nil {
			return nil, err
		}
		for index2, account := range result.Value {
			realIndex := index1*100 + index2
			mAccountInfo := &rpc.KeyedAccount{
				Pubkey:  accounts[realIndex],
				Account: account,
			}
			out = append(out, mAccountInfo)
		}
		fmt.Println(result.Context.Slot)
	}
	return out, nil
}
