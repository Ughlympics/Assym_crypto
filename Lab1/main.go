package main

import (
	"Lab1/gens"
	"fmt"
)

func main() {
	result := gens.LowLehmer(1, 20)
	fmt.Println("Result:", result)
	gens.Prob_test(result, 0.05)

	seq20 := gens.L20(12345, 30)
	fmt.Println("L20:", seq20)

	seq89 := gens.L89(6789346536456, 80)
	fmt.Println("L89:", seq89)
}
