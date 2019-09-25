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

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func Adder() {
	pos, neg := adder(), adder()
	fmt.Println(pos(10))
	fmt.Println(pos(10))
	fmt.Println(pos(10))
	fmt.Println(neg(2))
	//for i:= 0; i < 10; i++ {
	//	fmt.Println(
	//		pos(i),
	//		neg(-2*i))
	//}
}
