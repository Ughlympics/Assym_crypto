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

	seq20 := gens.L20(12345, 30)
	fmt.Println("L20:", seq20)

	seq89 := gens.L89(6789346536456, 80)
	fmt.Println("L89:", seq89)
}
