package pointers

import (
	"fmt"
)

type Vertex struct{
	X, Y float64
}

/*func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}*/

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Main() {
	v := Vertex{3,4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4,3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}