package main

import (
	"fmt"
	"github.com/zeroray/learning-golang/methods"
	"math"
)

func main() {
/*	v := methods.Vertex{3,4}
	fmt.Println(v.Abs())
*/

	f := methods.MyFloat(-math.Sqrt(2))
	fmt.Println(f.Abs())
}
