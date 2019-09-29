package main

import (
	"fmt"
	"github.com/zeroray/learning-golang/functions"
)

// https://en.wikipedia.org/wiki/Fibonacci_number

func main() {
	f := functions.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f(i))
	}
}
