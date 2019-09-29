package functions

var a []int

// fibonacci is a function that returns
// a function that return an int
func Fibonacci() func(int) int {
	return func(x int) int {
		if x >= 2 {
			xf := a[x-1] + a[x-2]
			a = append(a, xf)
			return xf
		}

		a = append(a, x)
		return (a[x])
	}
}


