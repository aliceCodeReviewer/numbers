package numbers

func Mul(n1, n2 int) int {
	return n1 * n2
}

func Mul3(n1, n2, n3 int) int {
	return Mul(Mul(n1, n2), n3)
}
