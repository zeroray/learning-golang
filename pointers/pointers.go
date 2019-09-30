package pointers

import (
	"fmt"
	"math"
)

type Vertex struct{
	X, Y float64
}

type User struct {
	name string
	age int
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Change u User to u *User and vice versa
func (u *User) changeName(name string) {
	u.name = name
}

// Change u User to u *User and vice versa
func (u *User) birtday() {
	u.age = u.age + 1
}

func Main() {
	u := User{"Daniel", 32}
	fmt.Println(u)
	u.birtday()
	fmt.Println(u)
	u.changeName("Daniel Fuentes")
	fmt.Println(u)

}