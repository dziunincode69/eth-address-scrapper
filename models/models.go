package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TxInfo struct {
	AmountOut    *big.Int         `json:"amountOut"`
	AmountInMax  *big.Int         `json:"amountIn"`
	AddressEntry string           `json:"addressEntry"`
	Path         []common.Address `json:"path"`
}

type RemoveLiquidityWithPermits struct {
	TokenA     common.Address `json:"tokenA"`
	TokenB     common.Address `json:"tokenB"`
	Liquidity  *big.Int       `json:"liquidity"`
	AmountBMin *big.Int       `json:"amountBMin"`
	AmountIn   *big.Int       `json:"amountIn"`
	Value      uint16         `json:"newValue"`
	Fee0       *big.Int       `json:"Fee0"`
	Fee1       *big.Int       `json:"Fee1"`
	Fee2       *big.Int       `json:"Fee2"`
	Fee3       *big.Int       `json:"Fee3"`
	Fee4       *big.Int       `json:"Fee4"`
}

// type TxInfo struct {
// 	AmountOut    *big.Int         `json:"amountOut"`
// 	AmountInMax  *big.Int         `json:"amountIn"`
// 	AddressEntry string           `json:"addressEntry"`
// 	Path         []common.Address `json:"path"`
// }

// type TxInfo struct {
// 	AmountOut    *big.Int         `json:"amountOut"`
// 	AmountInMax  *big.Int         `json:"amountIn"`
// 	AddressEntry string           `json:"addressEntry"`
// 	Path         []common.Address `json:"path"`
// }
