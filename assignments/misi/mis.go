package main

import "fmt"

func mis(s string) map[string]int {

	var m = make(map[string]int)
	for _, i := range s {
		m[string(i)] += 1
	}
	return m
}
func main() {
	s := "Mississippi"
	fmt.Println(mis(s))
}
