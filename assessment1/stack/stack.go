package main

import "fmt"

func push(s []int, x int) []int {
	s = append(s, x)
	return s
}

func pop(s []int) []int {
	s = s[:len(s)-1]
	return s

}
func main() {
	s := []int{3, 2, 5, 4}
	fmt.Println(s)
	s = push(s, 2)
	fmt.Println(s)
	s = push(s, 12)
	fmt.Println(s)
	s = pop(s)
	fmt.Println(s)
}
