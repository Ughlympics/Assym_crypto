package main

import (
	"Lab1/gens"
	"fmt"
)

func main() {
	result := gens.LowLehmer(1, 10)
	fmt.Println("Result:", result)

	for i := 1; i <= 5; i++ {
		fmt.Println("n =", i, " → ", gens.LowLehmer(1, uint64(i)))
	}
}
