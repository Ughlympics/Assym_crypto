package rsa

import (
	"math/big"
)

func Encrypt(message, e, n *big.Int) *big.Int {
	ciphertext := new(big.Int).Exp(message, e, n)
	return ciphertext
}
