package gens

import (
	"crypto/rand"
	"math/big"
)

// Blum-Micali values
var pStr1 = "CEA42B987C44FA642D80AD9F51F10457690DEF10C83D0BC1BCEE12FC3B6093E3"
var aStr1 = "5B88C41246790891C095E2878880342E88C79974303BD0400B090FE38A688356"

var p1, _ = new(big.Int).SetString(pStr1, 16)
var a1, _ = new(big.Int).SetString(aStr1, 16)

// Blum-Blum-Shub values
var pStr2 = "D5BBB96D30086EC484EBA3D7F9CAEB07"
var aStr2 = "425D2B9BFDB25B9CF6C416CC6E37B59C1F"

var p2, _ = new(big.Int).SetString(pStr2, 16)
var a2, _ = new(big.Int).SetString(aStr2, 16)
var m = new(big.Int).Mul(p2, a2)

func BMGenerator_byte(n int) ([]byte, error) {
	T0, err := rand.Int(rand.Reader, p1)
	if err != nil {
		return nil, err
	}
	if T0.Sign() == 0 {
		T0 = big.NewInt(1)
	}

	T := new(big.Int).Set(T0)
	bytes := make([]byte, n)

	interval := new(big.Int).Sub(p1, big.NewInt(1))
	interval.Div(interval, big.NewInt(256))

	for i := 0; i < n; i++ {

		k := new(big.Int).Div(T, interval)
		if k.Cmp(big.NewInt(256)) >= 0 {
			k = big.NewInt(255)
		}
		bytes[i] = byte(k.Int64())

		T.Exp(a1, T, p1)
	}

	return bytes, nil
}

func BMGenerator_bit(n int) ([]byte, error) {
	T0, err := rand.Int(rand.Reader, p1)
	if err != nil {
		return nil, err
	}
	if T0.Sign() == 0 {
		T0 = big.NewInt(1)
	}

	T := new(big.Int).Set(T0)
	result := make([]byte, n)

	for i := 0; i < n; i++ {
		var b byte = 0
		for bit := 0; bit < 8; bit++ {
			bitVal := T.Bit(0)
			b = (b << 1) | byte(bitVal)

			T.Exp(a1, T, p1)
		}
		result[i] = b
	}

	return result, nil
}

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

func BBSGenerator_bit(n int) ([]byte, error) {
	r0, err := rand.Int(rand.Reader, m)
	if err != nil {
		return nil, err
	}
	if r0.Sign() == 0 {
		r0 = big.NewInt(1)
	}

	r := new(big.Int).Set(r0)
	bytes := make([]byte, n)

	for i := 0; i < n; i++ {
		var b byte = 0
		for bit := 0; bit < 8; bit++ {
			r.Exp(r, big.NewInt(2), m)
			bitVal := r.Bit(0)
			b = (b << 1) | byte(bitVal)
		}
		bytes[i] = b
	}

	return bytes, nil
}
