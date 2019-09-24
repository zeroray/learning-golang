package variables

import "fmt"

// var x, y, z int
// var c, python, java bool

// var x, y, z int = 1, 2, 3
// var c, python, java = true, false, "No!"

func main() {
	// short declare
	var x, y, z int = 1, 2, 3
	c, python, java := true, false, "no!"
	fmt.Println(x, y, z, c, python, java)
}
