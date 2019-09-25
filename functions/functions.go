package functions

import (
	"fmt"
	"math"
)

func ValuesTypeFunction() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
}
