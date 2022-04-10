package prime

/*PrimeOrNot is used to check given number is prime or not*/
func PrimeOrNot(x int) string {
	for i := 2; i <= x/2; i++ {
		if x%i == 0 {
			return "false"
		}
	}
	return "true"
}

/*Prime generate list of primes in a given range*/
func Prime(n int) []int {
	var output []int
	for i := 2; i < n; i++ {
		x := PrimeOrNot(i)
		if x == "true" {
			output = append(output, i)
		}
	}
	return output
}
