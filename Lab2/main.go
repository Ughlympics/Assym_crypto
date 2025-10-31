package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("This is Lab2. Please refer to the assignment instructions.")
	var aStr2 = "425D2B9BFDB25B9CF6C416CC6E37B59C1F"
	var a2, _ = new(big.Int).SetString(aStr2, 16)
	fmt.Println("a2 byte length:", a2.BitLen())
	r0, err := rand.Int(rand.Reader, a2)
	if err != nil {
		fmt.Println("Error generating random number:", err)
		return
	}
	fmt.Println("Generated random number:", r0)
}
