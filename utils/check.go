package utils

import (
	"context"
	"time"

	"git/insiderScrapper/client"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func WaitTxConfirmed(ctx context.Context, hash common.Hash) (int, error) {
	for {
		tx, pending, err := client.Client().TransactionByHash(ctx, hash)
		if err != nil {
			// log.Print(err)
			return 99, err
		}
		if !pending {
			txStatus, _ := bind.WaitMined(ctx, client.Client(), tx)
			return int(txStatus.Status), nil
		}
		time.Sleep(time.Millisecond * 500)
	}

}
