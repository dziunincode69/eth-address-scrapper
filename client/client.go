package client

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
)

var (
	client *ethclient.Client
	rdb    *redis.Client

	err error
)

func ConnectClient(wsUrl string) {
	client, err = ethclient.Dial(wsUrl)
	if err != nil {
		log.Fatal(err)
	}
}
func ConnectRedis(password string) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: password,
		DB:       0,
	})
	fmt.Println("[✓] successfully connected to redis client")
}
func WriteConfig() {
	file, err := os.Create("config.yml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	confstring := `
MaxVictimBuy: 55500000000000000
MinVictimSell: 155000000000000000
ScrapToLatest: true
WebSocket: ws://127.0.0.1:8546`

	// Write some data to the file
	_, err = file.WriteString(confstring)
	if err != nil {
		panic(err)
	}
}

func Client() *ethclient.Client {
	return client
}
func Redis() *redis.Client {
	return rdb
}
