package gens

import (
	"crypto/rand"
	"math/big"
)

var pStr = "CEA42B987C44FA642D80AD9F51F10457690DEF10C83D0BC1BCEE12FC3B6093E3"
var aStr = "5B88C41246790891C095E2878880342E88C79974303BD0400B090FE38A688356"

var p, _ = new(big.Int).SetString(pStr, 16)
var a, _ = new(big.Int).SetString(aStr, 16)

func BMGenerator(n int) ([]byte, error) {
	T0, err := rand.Int(rand.Reader, p)
	if err != nil {
		return nil, err
	}
	if T0.Sign() == 0 {
		T0 = big.NewInt(1)
	}

	T := new(big.Int).Set(T0)
	bytes := make([]byte, n)

	interval := new(big.Int).Sub(p, big.NewInt(1))
	interval.Div(interval, big.NewInt(256))

	for i := 0; i < n; i++ {

		k := new(big.Int).Div(T, interval)
		if k.Cmp(big.NewInt(256)) >= 0 {
			k = big.NewInt(255)
		}
		bytes[i] = byte(k.Int64())

		T.Exp(a, T, p)
	}

	return bytes, nil
}
