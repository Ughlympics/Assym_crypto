package gens

func LowLehmer(x0 uint64, n uint64) []uint8 {
	var m = uint64(1 << 32)
	var a = uint64((1 << 16) + 1)
	var c = uint64(119)
	var x = x0

	result := make([]uint8, n)

	for i := uint64(0); i < n; i++ {
		x = (a*x + c) % m
		result[i] = uint8(x)
	}

	return result
}

func HighLehmer(x0 uint64, n uint64) []uint8 {
	var m = uint64(1 << 32)
	var a = uint64((1 << 16) + 1)
	var c = uint64(119)
	var x = x0

	result := make([]uint8, n)

	for i := uint64(0); i < n; i++ {
		x = (a*x + c) % m
		result[i] = uint8(x >> 24)
	}

	return result
}
