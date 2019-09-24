package exampleStruct

import "fmt"

type Vertex struct {
	X, Y int
}

func DeclareVertex() {
	//v := new(Vertex)
	var v *Vertex = new(Vertex)
	fmt.Println(v) // initial in 0
	v.X, v.Y = 11, 9
	fmt.Println(v)
	fmt.Println(*v) // Show value
}
