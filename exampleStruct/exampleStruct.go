package exampleStruct

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// var m map[string]Vertex
// var m = map[string]Vertex{
// 	"My Company": Vertex{
// 		33.333333, -77.7777777,
// 	},
// 	"Other Company": Vertex{
// 		44.444444, -55.000005,
// 	},
// }

var m = map[string]Vertex{
	"My Company":    {33.333333, -77.7777777},
	"Other Company": {44.444444, -55.000005},
}

func DeclareVertex() {
	// m = make(map[string]Vertex)
	// m["My Company"] = Vertex{
	// 	33.33333332, -77.777777,
	// }
	// fmt.Println(m["My Company"])
	fmt.Println(m)
}
