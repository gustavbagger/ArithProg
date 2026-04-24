package primality

import (
	"fmt"
	"math"

	pr "github.com/fxtlabs/primes"
)

func isPrimePow(val int64, n int) bool {
	fval := math.Pow(float64(val), 1.0/float64(n)) //might get some false positives here, since prod might not be nth pow
	fmt.Println(fval)
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
	}
	prod += 1

	return prod, isPrimePow(prod, n) //This should return prod,bool where bool := "prod is a prime power"
}
