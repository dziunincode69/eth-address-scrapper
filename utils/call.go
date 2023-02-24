package utils

import (
	"git/insiderScrapper/client"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	abis "git/insiderScrapper/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/liyue201/erc20-go/erc20"
)

func CheckWhitelistAddr(address string) {
	resp, err := http.Get("https://liongfamily.net/check?address=" + address)
	if err != nil {
		log.Fatal(err)
	}
	bodys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	body := string(bodys)
	if strings.Contains(body, "Hari") {
	} else {
		log.Fatal(body)
	}
	// return body
}
func CallName(token string) (string, string) {
	tokens := common.HexToAddress(token)
	tokensss, _ := erc20.NewGGToken(tokens, client.Client())

	name, _ := tokensss.Name(nil)
	symbol, _ := tokensss.Symbol(nil)
	return name, symbol
}
func GetPairs(contract string) common.Address {
	instance, err := abis.NewFactoryCaller(common.HexToAddress("0x5c69bee701ef814a2b6a3edd4b1652cb9cc5aa6f"), client.Client())
	if err != nil {
		log.Fatal(err)
	}
	call := &bind.CallOpts{}
	contracts := common.HexToAddress(contract)
	GetPair, err := instance.GetPair(call, common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"), contracts)
	if err != nil {
		log.Print("Error Pair call", err)
	}
	return GetPair

}
