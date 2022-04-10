package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v, "hello")
	default:
		fmt.Println("byee")
	}
}
func main() {
	do(5)
	do("hello")

}
