package slices

import (
	"fmt"
)

func ReadSlice() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}

func SliceOfSlice() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])

	// missing low index implies 0
	fmt.Println("p[:3] ==", p[:3])

	// missing high index implues len(s)
	fmt.Println("p[4:] ==", p[4:])
}

func CreateSlices() {
	a := make([]int, 5, 10)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
	// more long full capacity
	z := a[:10]
	printSlice("z", z)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func SliceNil() {
	var z []int
	printSlice("z", z)
	if z == nil {
		fmt.Println("Nil!")
	}
}


