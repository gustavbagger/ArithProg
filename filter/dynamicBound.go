package filter

import (
	"math"
)

// 1.7 works, 1.4193747081548531 works if g(p)<p^0.621526 (since h = 0.82 * p^1/4 is admissible)
var C float64 = 1.4193747081548531

var PrimeListUpperBound int = 1 << 20

func DeltaSum(list []int, gContribution float64) float64 {
	var sum float64
	for _, p := range list {
		sum -= 1.0 / float64(p)
	}
	return sum + 1 - gContribution
}

func Triangle(m, r, s int, delta float64) float64 {
	return float64(2+m*r+s-1) / delta
}

func InitBestR(omegaMax, s int, primeList []int, gContribution, gProd float64) []int {
	bestS := make([]int, omegaMax+1)
	for omega := 1; omega <= omegaMax; omega++ {
		sBest := 0
		currentBest := 4.0 * Triangle(4, 0, s, 1.0-gContribution) * float64(int(1)<<(omega+1))
		for r := 1; r <= omega; r++ {
			delta := DeltaSum(primeList[omega-r:omega], gContribution)
			if delta <= 0.0 {
				break
			}
			currentTriangle := Triangle(4, r, s, delta)
			currentTry := 4.0 * currentTriangle * gProd * float64(int(1)<<(4*(omega-r)))

			if currentTry < currentBest {
				currentBest = currentTry

				sBest = s

			}
		}
		bestS[omega] = sBest
	}
	return bestS
}

func OptSieveBoundLog(omega, r, s, n int, indexes, primeList []int, gContribution, gProd, boundLog float64) float64 {
	if r == 0 || len(indexes) < r {
		return boundLog
	}
	return math.Min(boundLog, PSieveLog(omega, r, s, n, indexes, primeList, gContribution, gProd))
}

func PSieveLog(omega, r, s, n int, indexes, primeList []int, gContribution, gProd float64) float64 {
	last := make([]int, r)
	for i := 0; i < r; i++ {
		last[i] = primeList[indexes[omega-r+i]]
	}
	delta := DeltaSum(last, gContribution)
	if delta <= 0.0 {
		return 0.0
	}

	return 2.0 *
		(math.Log(4) +
			math.Log(Triangle(4, r, s, delta)) +
			math.Log(gProd) +
			4.0*float64(omega-r)*math.Log(2.0))

}
