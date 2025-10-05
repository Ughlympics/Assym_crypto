package gens

import (
	"fmt"
	"math"
)

func Prob_test(data []byte, alpha float64) {
	m := len(data)
	if m == 0 {
		fmt.Println("freq test: no data")
		return
	}

	//  v_j
	counts := make([]float64, 256)
	for _, b := range data {
		counts[int(b)]++
	}

	// hi^2
	expected := float64(m) / 256.0
	var chiStat float64
	for _, v := range counts {
		diff := v - expected
		chiStat += diff * diff / expected
	}

	// crit hi^2
	k := 255.0
	quantile := ChiSquareQuantile(1.0-alpha, k)

	pass := chiStat <= quantile
	fmt.Printf("hi^2 stat = %.6f\n", chiStat)
	fmt.Printf("crit value (α=%.4f, k=%.0f) ≈ %.6f\n", alpha, k, quantile)
	fmt.Printf("hypothesis H0 accepted: %v\n", pass)
}

func ChiSquareQuantile(p, k float64) float64 {
	if p <= 0 {
		return 0
	}
	if p >= 1 {
		return math.Inf(1)
	}
	z := NormInv(p)
	term := 1.0 - 2.0/(9.0*k) + z*math.Sqrt(2.0/(9.0*k))
	return k * term * term * term
}

func NormInv(p float64) float64 {

	a1 := -3.969683028665376e+01
	a2 := 2.209460984245205e+02
	a3 := -2.759285104469687e+02
	a4 := 1.383577518672690e+02
	a5 := -3.066479806614716e+01
	a6 := 2.506628277459239e+00

	b1 := -5.447609879822406e+01
	b2 := 1.615858368580409e+02
	b3 := -1.556989798598866e+02
	b4 := 6.680131188771972e+01
	b5 := -1.328068155288572e+01

	c1 := -7.784894002430293e-03
	c2 := -3.223964580411365e-01
	c3 := -2.400758277161838e+00
	c4 := -2.549732539343734e+00
	c5 := 4.374664141464968e+00
	c6 := 2.938163982698783e+00

	d1 := 7.784695709041462e-03
	d2 := 3.224671290700398e-01
	d3 := 2.445134137142996e+00
	d4 := 3.754408661907416e+00

	// Define break-points.
	pLow := 0.02425
	pHigh := 1 - pLow

	var q, r float64
	if p < 0 || p > 1 {
		return math.NaN()
	} else if p == 0 {
		return math.Inf(-1)
	} else if p == 1 {
		return math.Inf(1)
	} else if p < pLow {
		// Rational approximation for lower region
		q = math.Sqrt(-2 * math.Log(p))
		return (((((c1*q+c2)*q+c3)*q+c4)*q+c5)*q + c6) /
			(((d1*q+d2)*q+d3)*q + d4)
	} else if p <= pHigh {
		// Rational approximation for central region
		q = p - 0.5
		r = q * q
		return (((((a1*r+a2)*r+a3)*r+a4)*r+a5)*r + a6) * q /
			(((((b1*r+b2)*r+b3)*r+b4)*r+b5)*r + 1.0)
	} else {
		// Rational approximation for upper region
		q = math.Sqrt(-2 * math.Log(1-p))
		return -(((((c1*q+c2)*q+c3)*q+c4)*q+c5)*q + c6) /
			(((d1*q+d2)*q+d3)*q + d4)
	}
}

func Ind_test(data []byte, alpha float64) {
	m := len(data)
	if m < 2 {
		fmt.Println("Помилка: послідовність занадто коротка")
		return
	}
	nPairs := m / 2
	if nPairs == 0 {
		fmt.Println("Помилка: немає пар для аналізу")
		return
	}

	var counts [256][256]int64
	var rowSums, colSums [256]int64

	for i := 0; i < nPairs; i++ {
		b1 := data[2*i]
		b2 := data[2*i+1]
		counts[b1][b2]++
		rowSums[b1]++
		colSums[b2]++
	}

	chi := 0.0
	n := float64(nPairs)
	minExpected := math.Inf(1)
	zeroExpectedCells := 0

	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			expected := (float64(rowSums[i]) * float64(colSums[j])) / n
			if expected == 0 {
				zeroExpectedCells++
				continue
			}
			if expected < minExpected {
				minExpected = expected
			}
			diff := float64(counts[i][j]) - expected
			chi += diff * diff / expected
		}
	}

	df := float64(255 * 255)
	quant := ChiSquareQuantile(1.0-alpha, df)
	pass := chi <= quant

	fmt.Printf("χ² = %.6f\n", chi)
	fmt.Printf("Crit value (α=%.4f, df=%d) ≈ %.6f\n", alpha, 255*255, quant)

	if pass {
		fmt.Println("Independence hypothesis accepted.")
	} else {
		fmt.Println("Independence hypothesis rejected.")
	}
}
