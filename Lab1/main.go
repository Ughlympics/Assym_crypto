package main

import (
	"Lab1/gens"
	"crypto/rand"
	"fmt"
)

func main() {
	var r = 10
	var alpha = 0.1
	var testBytes uint64 = 500000
	fmt.Println("///////////////////////////////////////////////////////////////////////////")
	fmt.Println("All tests made at a length of 500,000 bytes and on choosen a(will be print):")
	fmt.Println("///////////////////////////////////////////////////////////////////////////")

	fmt.Println("Built in generator quality tests:")
	data1 := make([]byte, testBytes)

	_, err := rand.Read(data1)
	if err != nil {
		panic(err)
	}
	gens.QualityTest(data1, alpha, r)

	fmt.Println("///////////////////////////////////////////////////////////////////////////")
	fmt.Println("LowLehmer generator quality tests:")
	data2 := gens.LowLehmer(1, testBytes)
	gens.QualityTest(data2, alpha, r)

	fmt.Println("///////////////////////////////////////////////////////////////////////////")
	fmt.Println("HighLehmer generator quality tests:")
	data3 := gens.HighLehmer(1, 1000000)
	gens.QualityTest(data3, alpha, r)

	fmt.Println("///////////////////////////////////////////////////////////////////////////")
	fmt.Println("L20 generator quality tests:")
	//data4 := gens.L20(12345, 500000)
	//gens.QualityTest(data4, alpha, r)
	seedL20 := []bool{true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false}
	checkBits := 800000
	data4 := gens.GenerateL20(seedL20, 800000)

	gens.QualityTest(data4, alpha, r)

	fmt.Println("///////////////////////////////////////////////////////////////////////////")
	fmt.Println("L89 generator quality tests:")
	//data5 := gens.L89(6789346536456, 500000)
	//gens.QualityTest(data5, alpha, r)
	seedL89 := []bool{true, true, false, true, false, true, false, false, true, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false}
	data5 := gens.GenerateL89(seedL89, checkBits)
	gens.QualityTest(data5, alpha, r)

	fmt.Println("///////////////////////////////////////////////////////////////////////////")
	fmt.Println("Giffi generator quality tests:")
	seedL11 := []bool{true, false, true, false, true, false, true, false, true, false, true}
	seedL9 := []bool{true, true, false, true, false, true, false, false, true}
	seedL10 := []bool{true, false, false, true, true, false, true, false, false, true}

	//checkBits := 8000000
	x := gens.GenerateL11(seedL11, checkBits)
	y := gens.GenerateL9(seedL9, checkBits)
	s := gens.GenerateL10(seedL10, checkBits)

	zBits := make([]bool, checkBits)
	for i := 0; i < checkBits; i++ {
		zBits[i] = gens.F(x[i], y[i], s[i])
	}

	data6 := gens.BitsToBytes(zBits)
	gens.QualityTest(data6, alpha, r)

	fmt.Println("///////////////////////////////////////////////////////////////////////////")
	fmt.Println("Wolfram generator quality tests:")
	bits1 := gens.WolframGenerator(0xDEADBEEF, 8000000)
	data7 := gens.BitsToBytes(bits1)
	gens.QualityTest(data7, alpha, r)

	fmt.Println("///////////////////////////////////////////////////////////////////////////")
	fmt.Println("Librarian generator quality tests:")
	data8, err := gens.LibrarianGenerator("Noviy_zavet_f.txt", 500000)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	gens.QualityTest(data8, alpha, r)

	fmt.Println("///////////////////////////////////////////////////////////////////////////")
	fmt.Println("BM bit generator quality tests:")
	data9, err := gens.BMGenerator_bit(500000)
	if err != nil {
		panic(err)
	}
	gens.QualityTest(data9, alpha, r)

	fmt.Println("///////////////////////////////////////////////////////////////////////////")
	fmt.Println("BM byte generator quality tests:")
	data10, err := gens.BMGenerator_byte(500000)
	if err != nil {
		panic(err)
	}
	gens.QualityTest(data10, alpha, r)

	fmt.Println("///////////////////////////////////////////////////////////////////////////")
	fmt.Println("BBS bit generator quality tests:")
	data11, err2 := gens.BBSGenerator_bit(500000)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}
	gens.QualityTest(data11, alpha, r)

	fmt.Println("///////////////////////////////////////////////////////////////////////////")
	fmt.Println("BBS byte generator quality tests:")
	data12, err3 := gens.BBSGenerator_byte(500000)
	if err3 != nil {
		fmt.Println("Error:", err3)
		return
	}
	gens.QualityTest(data12, alpha, r)
}
