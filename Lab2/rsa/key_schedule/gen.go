package rsa

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var pStr2 = "D5BBB96D30086EC484EBA3D7F9CAEB07"
var aStr2 = "425D2B9BFDB25B9CF6C416CC6E37B59C1F"

var p2, _ = new(big.Int).SetString(pStr2, 16)
var a2, _ = new(big.Int).SetString(aStr2, 16)
var m = new(big.Int).Mul(p2, a2)

func BBSGenerator_byte(n int) (*big.Int, error) {
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

	result := new(big.Int).SetBytes(bytes)
	return result, nil
}

func MillerRabinTest(p *big.Int, k int) (bool, error) {
	for i := 0; i < k; i++ {
		x, ok := step1(p)
		if !ok {
			fmt.Println("Found a non-trivial factor:", x)
			return false, nil
		}
		if !step2(p, x) {
			return false, nil
		}
	}
	return true, nil
}

func decomposition(n *big.Int) (d *big.Int, r *big.Int) {
	d = new(big.Int).Set(n)
	r = big.NewInt(0)
	for i := 0; ; i++ {
		if new(big.Int).Mod(d, big.NewInt(2)).Cmp(big.NewInt(0)) != 0 {
			return d, r
		}
		d.Div(d, big.NewInt(2))
		r.Add(r, big.NewInt(1))
	}
}

func step1(n *big.Int) (*big.Int, bool) {
	randNum, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, false
	}

	g := new(big.Int).GCD(nil, nil, randNum, n)
	if g.Cmp(big.NewInt(1)) != 0 {
		return g, false
	}
	return randNum, true
}

func step2(p, a *big.Int) bool {
	one := big.NewInt(1)
	minusOne := new(big.Int).Sub(p, one)

	// p-1 = 2^s * d
	d, s := decomposition(new(big.Int).Sub(p, one))

	// x = a^d mod p
	x := new(big.Int).Exp(a, d, p)
	if x.Cmp(one) == 0 || x.Cmp(minusOne) == 0 {
		return true
	}

	for i := int64(1); i < s.Int64(); i++ {
		x.Exp(x, big.NewInt(2), p)

		if x.Cmp(minusOne) == 0 {
			return true
		}
		if x.Cmp(one) == 0 {
			return false
		}
	}
	return false
}
