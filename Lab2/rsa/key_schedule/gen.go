package rsa

import (
	"crypto/rand"
	"math/big"
)

var pStr2 = "D5BBB96D30086EC484EBA3D7F9CAEB07"
var aStr2 = "425D2B9BFDB25B9CF6C416CC6E37B59C1F"

var p2, _ = new(big.Int).SetString(pStr2, 16)
var a2, _ = new(big.Int).SetString(aStr2, 16)
var m = new(big.Int).Mul(p2, a2)

func BBSGenerator_byte(n int) ([]byte, error) {
	r0, err := rand.Int(rand.Reader, p2)
	if err != nil {
		return nil, err
	}
	if r0.Sign() == 0 {
		r0 = big.NewInt(1)
	}

	r := new(big.Int).Set(r0)
	bytes := make([]byte, n)

	for i := 0; i < n; i++ {
		r.Exp(r, big.NewInt(2), m)
		mod256 := new(big.Int).Mod(r, big.NewInt(256))
		bytes[i] = byte(mod256.Int64())
	}

	return bytes, nil
}
