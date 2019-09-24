package exampleStruct

import "fmt"

type Vertex struct {
	X int
	Y int
}

func DeclareVertex() {
	v := Vertex{1, 3}
	v.X = 4
	fmt.Println(v.X)
}
