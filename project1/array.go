package main

import "fmt"

func array() {
	s := []int{1, 2, 3, 4, 5}
	k := &s
	s = s[0:3]
	j := &s
	s[0] = 12
	fmt.Println(s, k, j)
}
