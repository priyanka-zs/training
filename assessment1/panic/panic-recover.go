package main

import "fmt"

func main() {
	fmt.Println("first")
	defer fmt.Println("last")
	defer func() {
		if recover() != nil {
			fmt.Println("recovered from panic")
		}
	}()
	panic("panic occurred")
	fmt.Println("this will not be printed")

}
