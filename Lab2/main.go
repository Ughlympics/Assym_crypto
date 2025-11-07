package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	rsa "rsa/rsa/key_schedule"
)

func main() {
	var aStr2 = "E3580DB782ED794B5B3CA47348DE8B15"
	var a2, _ = new(big.Int).SetString(aStr2, 16)
	prime, _ := rand.Prime(rand.Reader, 512)
	fmt.Println("Generated prime:", prime)

	t1, _ := rsa.MillerRabinTest(a2, 10)
	t2, _ := rsa.MillerRabinTest(prime, 10)
	fmt.Println("Miller-Rabin test result:", t1, t2)
}
