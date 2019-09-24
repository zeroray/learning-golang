package exampleStruct

import "fmt"

type Vertex struct {
	X int
	Y int
}

func DeclareVertex() {
	p := Vertex{1, 2}
	q := &p
	q.X = 1e2
	fmt.Println(p)
}
