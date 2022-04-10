package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 1, 2, 3}
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	fmt.Println(s)
}
