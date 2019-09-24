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
