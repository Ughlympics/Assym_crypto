package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	features "rsa/rsa"
	key "rsa/rsa/key_schedule"
)

func main() {
	var aStr2 = "E3580DB782ED794B5B3CA47348DE8B15"
	var a2, _ = new(big.Int).SetString(aStr2, 16)
	prime, _ := rand.Prime(rand.Reader, 512)
	fmt.Println("Generated prime:", prime)

	t1, _ := key.MillerRabinTest(a2, 10)
	t2, _ := key.MillerRabinTest(prime, 10)
	fmt.Println("Miller-Rabin test result:", t1, t2)

	p, q, _ := key.GenKey(32)
	fmt.Printf("p = 0x%X\nq = 0x%X\n", p, q)

	n := new(big.Int).Mul(p, q)

	e := big.NewInt(65537)
	message := big.NewInt(123456789)

	ciphertext := features.Encrypt(message, e, n)

	fmt.Printf("n = 0x%X\n", n)
	fmt.Printf("Ciphertext = 0x%X\n", ciphertext)
}
