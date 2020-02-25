package main

import "fmt"

func initConditionPost() {
	// for init; condition; post {}

	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j <= 3; j++ {
			fmt.Printf("\t\tThe inner loop: %d\n", j)
		}
	}
}

func onlyFor() {
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")
}

func breakContinue() {
	x := 0
	for {
		x++
		fmt.Println(x)
		if x > 100 {
			break
		}

		if x%2 != 0 {
			continue
		}

		fmt.Println(x)
	}
	fmt.Println("done.")
}

func main() {
	// initConditionPost()
	// onlyFor()
	breakContinue()
}
