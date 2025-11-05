package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	rsa "rsa/rsa/key_schedule"
)

func main() {
	fmt.Println("This is Lab2. Please refer to the assignment instructions.")
	var aStr2 = "35"
	var a2, _ = new(big.Int).SetString(aStr2, 16)
	fmt.Println("a2 byte length:", a2.BitLen())
	r0, err := rand.Int(rand.Reader, a2)
	if err != nil {
		fmt.Println("Error generating random number:", err)
		return
	}
	t, _ := rsa.MillerRabinTest(a2, 10)
	fmt.Println("Generated random number:", r0)
	fmt.Println("Miller-Rabin test result:", t)
}
