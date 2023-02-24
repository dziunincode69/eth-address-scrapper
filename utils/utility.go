package utils

import (
	"fmt"
	"git/insiderScrapper/client"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/params"
	"github.com/spf13/viper"
)

func WeiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}
func SaveToFile(s string) {
	file, err := os.OpenFile("rez.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the string to the file
	str := s
	_, err = file.WriteString(str)
	if err != nil {
		panic(err)
	}
}
func InitializeVipers() {
	//encode key in bytes to string and keep as secret, put in a vault

	viper.SetConfigFile("config.yml")
	if err := viper.ReadInConfig(); err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			client.WriteConfig()
			fmt.Println("Config File Has created")
			os.Exit(1)
		}
	}
}
