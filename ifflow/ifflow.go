package ifflow

import (
	"fmt"
	"math"
)

func Sqrt(x float64) string {
	if x < 0 {
		return Sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func Pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
