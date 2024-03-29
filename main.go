package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"git/insiderScrapper/client"
	"git/insiderScrapper/utils"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"
)

var fromBlock int
var toBlock int
var listBuy string
var listSell string

func main() {
	utils.InitializeVipers()
	client.ConnectClient(viper.GetString("WebSocket"))
	fmt.Print("From Block: ")
	fmt.Scanln(&fromBlock)
	if !viper.GetBool("ScrapToLatest") {
		fmt.Print("To Block: ")
		fmt.Scanln(&toBlock)
	} else {
		toBlocks, err := client.Client().HeaderByNumber(context.Background(), nil)
		toBlock = int(toBlocks.Number.Int64())
		if err != nil {
			fmt.Println("Block eRROR", err)
		}
	}
	chainID, err := client.Client().NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfull connected to network chainId: " + chainID.String())
	var tokenAddress string
	for ii := fromBlock; ii < toBlock; ii++ {
		blockNumber := big.NewInt(int64(ii))
		fmt.Println("Block number: ", blockNumber)
		block, err := client.Client().BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			log.Fatal("Block error: ", err)
		}

		for _, tx := range block.Transactions() {
			txHash := tx.Hash().Hex()
			input := hex.EncodeToString(tx.Data())
			asMsg, _ := tx.AsMessage(types.NewLondonSigner(chainID), nil)
			from := asMsg.From()
			if strings.Contains(input, "fb3bdb41") || strings.Contains(input, "5ae401dc") && strings.Contains(input, "42712a67") {
				txReceipt, err := client.Client().TransactionReceipt(context.Background(), common.HexToHash(txHash))
				var ethUsed *big.Int
				var tokenYougot *big.Int
				if err != nil {
					fmt.Println(err)
				}
				for i := 0; i < len(txReceipt.Logs); i++ {
					interact := txReceipt.Logs[i].Address
					Datas := txReceipt.Logs[i].Data
					Topics := txReceipt.Logs[i].Topics
					if i == 2 {
						tokenAddress = interact.String()
					}
					// fmt.Println("Datas: ", hex.EncodeToString(Datas))
					// fmt.Println("Topics: ", Topics)
					if Topics == nil {
						fmt.Println()
					} else if Topics[0] == common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822") {
						decs := utils.SwapDecode(hex.EncodeToString(Datas), Datas)
						if strings.Contains(input, "5ae401dc") && strings.Contains(input, "42712a67") {
							ethUsed = decs.Fee1
							tokenYougot = decs.Fee2
						} else {
							ethUsed = decs.Fee0
							tokenYougot = decs.Fee3
						}
						name, sym := utils.CallName(tokenAddress)
						if ethUsed.Int64() != 0 && tokenYougot.Int64() != 0 {
							if ethUsed.Uint64() <= viper.GetUint64("MaxVictimBuy") {
								fmt.Println("========== B U Y ==========")
								fmt.Println("From", from)
								fmt.Println("Name:", name, "Symbol:", sym)
								fmt.Println("TokenAddress", tokenAddress)
								fmt.Println("Total WETH used", utils.WeiToEther(ethUsed), "ETH")
								fmt.Println("AmountOut", tokenYougot)
								fmt.Println("Tx Hash", txReceipt.Logs[i].TxHash)
								fmt.Println()
								listBuy += strings.ToLower(from.String() + ":BUY:" + strings.ToLower(tokenAddress) + ":" + utils.WeiToEther(ethUsed).String() + ":" + txReceipt.Logs[i].TxHash.String() + "\n")
							}
						}
					}
				}
			} else if strings.Contains(input, "791ac947") || strings.Contains(input, "18cbafe5") || strings.Contains(input, "3593564c") && strings.Contains(input, "80c") || strings.Contains(input, "5ae401dc") && strings.Contains(input, "472b43f3") && strings.Contains(input, "49404b7c") {
				txReceipt, err := client.Client().TransactionReceipt(context.Background(), common.HexToHash(txHash))
				var totalTokenOut *big.Int
				var TotalEthIn *big.Int
				var sign string
				if err != nil {
					fmt.Println(err)
				}
				for i := 0; i < len(txReceipt.Logs); i++ {
					interact := txReceipt.Logs[i].Address
					Datas := txReceipt.Logs[i].Data
					Topics := txReceipt.Logs[i].Topics
					if strings.Contains(input, "3593564c") && strings.Contains(input, "80c") { // execute
						if i == 1 {
							tokenAddress = interact.String()
						}
					} else if strings.Contains(input, "5ae401dc") && strings.Contains(input, "472b43f3") && strings.Contains(input, "49404b7c") {
						if i == 0 {
							tokenAddress = interact.String()
						}
					} else {
						if i == 0 {
							tokenAddress = interact.String()
						}
					}

					if Topics == nil {
						fmt.Println()
					} else if Topics[0] == common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822") {
						decs := utils.SwapDecode(hex.EncodeToString(Datas), Datas)
						if strings.Contains(input, "5ae401dc") && strings.Contains(input, "472b43f3") && strings.Contains(input, "49404b7c") {
							totalTokenOut = decs.Fee0
							TotalEthIn = decs.Fee3
						} else {
							totalTokenOut = decs.Fee1
							TotalEthIn = decs.Fee2
						}

						name, sym := utils.CallName(tokenAddress)
						if strings.Contains(input, "3593564c") && strings.Contains(input, "80c") {
							sign = "execute"
						} else if strings.Contains(input, "5ae401dc") && strings.Contains(input, "472b43f3") && strings.Contains(input, "49404b7c") {
							sign = "multicall"
						} else {
							sign = "SwapExactToken"
						}

						if TotalEthIn.Int64() != 0 && totalTokenOut.Int64() != 0 {
							if TotalEthIn.Uint64() >= viper.GetUint64("MinVictimSell") {
								regg := strings.ToLower(from.String()) + ":buy:" + strings.ToLower(tokenAddress)
								bal, _ := client.Client().BalanceAt(context.Background(), from, nil)
								if strings.Contains(listBuy, regg) {
									fmt.Println("========== I N S I D E R ==========")
									fmt.Println("From", from)
									fmt.Println("Name:", name, "| Symbol:", sym)
									fmt.Println("TokenAddress", tokenAddress)
									fmt.Println("Total Token Out", totalTokenOut)
									fmt.Println("Total WETH In", utils.WeiToEther(TotalEthIn), "ETH")
									fmt.Println("Tx Hash", txReceipt.Logs[i].TxHash)
									fmt.Println("Methode", sign)
									fmt.Println()
									utils.SaveToFile("INSIDER:" + from.String() + ":" + tokenAddress + ":" + name + ":" + ":" + sym + ":" + utils.WeiToEther(TotalEthIn).String() + "WETH:" + txReceipt.Logs[i].TxHash.String() + ":" + utils.WeiToEther(bal).String() + ":" + sign + "\n")
								}
								fmt.Println("========== S E L L ==========")
								fmt.Println("From", from)
								fmt.Println("Name:", name, "| Symbol:", sym)
								fmt.Println("TokenAddress", tokenAddress)
								fmt.Println("Total Token Out", totalTokenOut)
								fmt.Println("Total WETH In", utils.WeiToEther(TotalEthIn), "ETH")
								fmt.Println("Tx Hash", txReceipt.Logs[i].TxHash)
								fmt.Println("Methode", sign)
								fmt.Println()
							}
						}
					}
				}
			}

		}
	}
}
