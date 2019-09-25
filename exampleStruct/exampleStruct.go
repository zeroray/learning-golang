package exampleStruct

import "fmt"

func DeclareVertex() {
	m := make(map[string]int)

	m["Response"] = 42
	fmt.Println("The value:", m["Response"])

	m["Response"] = 48
	fmt.Println("The value:", m["Response"])

	delete(m, "Response")
	fmt.Println("The value:", m["Response"])

	v, ok := m["Response"]
	fmt.Println("The value:", v, "Present?", ok)

}
