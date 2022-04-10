package main

import "fmt"

func d() {
	n := 345647
	c := 0
	for n > 0 {
		c = c + 1
		n = n / 10
	}
	fmt.Printf("%v", 100)
}
