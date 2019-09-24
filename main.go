package main

import (
	"fmt"

	"github.com/zeroray/learning-golang/ifflow"
)

func main() {
	fmt.Println(
		ifflow.Pow(3, 2, 10),
		ifflow.Pow(3, 3, 20),
	)
}
