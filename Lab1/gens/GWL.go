package gens

func GenerateL11(seed []bool, n int) []bool {
	x := append([]bool(nil), seed...)
	for i := len(seed); i < n; i++ {
		newBit := x[i-11] != x[i-2]
		x = append(x, newBit)
	}
	return x
}

func GenerateL9(seed []bool, n int) []bool {
	y := append([]bool(nil), seed...)
	for i := len(seed); i < n; i++ {
		newBit := y[i-9] != y[i-1] != y[i-3] != y[i-4]
		y = append(y, newBit)
	}
	return y
}

func GenerateL10(seed []bool, n int) []bool {
	s := append([]bool(nil), seed...)
	for i := len(seed); i < n; i++ {
		newBit := s[i-10] != s[i-3]
		s = append(s, newBit)
	}
	return s
}

func F(x, y, s bool) bool {
	if s {
		return x
	}
	return y
}

func BitsToBytes(bits []bool) []byte {
	n := (len(bits) + 7) / 8
	bytes := make([]byte, n)
	for i, b := range bits {
		if b {
			bytes[i/8] |= 1 << uint(7-(i%8))
		}
	}
	return bytes
}
