package numbers

func Mul(n1, n2 int) int {
	out := 0
	for i := 0; i < n1; i++ {
		out = Add(out, n2)
	}
	return out
}

func Mul3(n1, n2, n3 int) int {
	return Mul(Mul(n1, n2), n3)
}
