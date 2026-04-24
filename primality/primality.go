package primality

import (
	"math"

	pr "github.com/fxtlabs/primes"
)

func isPrimePow(val int64, n int) bool {
	fval := math.Pow(float64(val), 1.0/float64(n)) //might get some false positives here, since prod might not be nth pow

	for exp := 1.0; exp <= math.Log2(fval); exp += 1.0 {
		testVal := int(math.Floor(math.Pow(fval, 1.0/exp)))
		if pr.IsPrime(testVal) {
			return true
		}
	}
	return false
}

func ValidExponentSet(indexes, exponents, allValues []int, n int) (int64, bool) {
	prod := int64(1)
	for i, index := range indexes {
		for range exponents[i] {
			prod *= int64(allValues[index])
		}
		if prod < 0 {
			return 0, false
		}
	}
	prod += 1

	return prod, true //This should return prod,bool where bool := "prod is a prime power"
}

func ValidExponentSet192(indexes, exponents, allValues []int, n int) (uint192, bool) {
	prod := uint192{Lo: 1}
	for i, index := range indexes {
		for exp := 1; exp <= exponents[i]; exp++ {
			prod = mulMod192(prod, uint192{Lo: uint64(allValues[index])})
		}
	}
	prod = add192(prod, uint192{Lo: 1})
	// prp := strongPRP(prod)
	return prod, true //This should return prod,bool where bool := "prod is a prime power"
}
