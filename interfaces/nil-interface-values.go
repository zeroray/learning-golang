package interfaces

type I interface {
	M()
}

func Main5() {
	var i I
	describe(i)
	i.M()
}
