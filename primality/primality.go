package primality

func ValidExponentSet(indexes, exponents, allValues []int) (uint64, bool) {
	prod := uint64(1)
	for i, index := range indexes {
		for range exponents[i] {
			prod *= uint64(allValues[index])
		}
	}
	prod += 1
	return prod, true //This should return prod,bool where bool := "prod is a prime power"
}
