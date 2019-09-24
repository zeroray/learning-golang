package ciclofor

import "fmt"

// SumNumber is a Public method
func SumNumber() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	showValues(sum)
}

// showValues is a private method
func showValues(value int) {
	fmt.Println(value)
}

// SumNumber1000 example pre and post sentences empty
func SumNumber1000() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	showValues(sum)
}

func Infinite() {
	var i int = 0
	for {
		i = i + 1
		showValues(i)
	}
}
