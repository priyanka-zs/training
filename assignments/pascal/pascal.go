package main

import "fmt"

//Pascal takes n as input and produces pascals array
func Pascal(n int) [][]int {
	a := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		a[i] = make([]int, i+1)
		a[i][0], a[i][i] = 1, 1
		for j := 1; j < i; j++ {
			a[i][j] = a[i-1][j] + a[i-1][j-1]
		}
	}
	return a
}
func main() {
	fmt.Println(Pascal(3))
}
