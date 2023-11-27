package numbers

func Add(number1, number2 int) int {
	return number1 + number2
}

func Div(number1, number2 int) int {
	return number1 - number2
}

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

func Add3(n1 int, n2 int, n3 int) int {
	return n1 + n2 + n3
}

func Pow(n1, n2 int) int {
	out := 1
	for i := 0; i < n1; i++ {
		out = Mul(out, n2)
	}
	return out
}
