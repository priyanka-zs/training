package main

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}

func compute(a int, b int, fn func(int, int) int) int {
	return fn(a, b)

}

/*func calci(a int, b int, s string) int {

	if s == "+" {
		return a + b
	} else if s == "-" {
		return a - b
	} else if s == "*" {
		return a * b
	} else if s == "/" {
		return a / b
	} else {
		return a % b
	}
}*/
