package exampleStruct

import "fmt"

type Vertex struct {
	X int
	Y int
}

func DeclareVertex() {
	fmt.Println(Vertex{1, 2})
}
