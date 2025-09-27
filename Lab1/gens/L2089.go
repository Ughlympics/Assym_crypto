package gens

func seedToBits(seed uint64, size int) []uint8 {
	bits := make([]uint8, size)
	for i := 0; i < size; i++ {
		if (seed>>i)&1 == 1 {
			bits[i] = 1
		} else {
			bits[i] = 0
		}
	}
	allZero := true
	for _, b := range bits {
		if b == 1 {
			allZero = false
			break
		}
	}
	if allZero {
		bits[0] = 1
	}
	return bits
}

// L20: x_t = x_{t-3} ⊕ x_{t-5} ⊕ x_{t-9} ⊕ x_{t-20}
func L20(seed uint64, n int) []uint8 {
	state := seedToBits(seed, 20)
	result := make([]uint8, n)

	for i := 0; i < n; i++ {
		result[i] = state[19]

		newBit := state[17] ^ state[15] ^ state[11] ^ state[0]

		for j := 19; j > 0; j-- {
			state[j] = state[j-1]
		}
		state[0] = newBit
	}
	return result
}

// L89: x_t = x_{t-38} ⊕ x_{t-89}
func L89(seed uint64, n int) []uint8 {
	state := seedToBits(seed, 89)
	result := make([]uint8, n)

	for i := 0; i < n; i++ {
		result[i] = state[88]

		newBit := state[51] ^ state[0]

		for j := 88; j > 0; j-- {
			state[j] = state[j-1]
		}
		state[0] = newBit
	}
	return result
}
