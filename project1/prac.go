package main

import "fmt"

func prac() {

	map_1 := map[int]string{
		1: "sun",
		2: "mon",
		3: "tues"}
	map_1[4] = "wed"
	b := map_1
	fmt.Println(map_1, b)
	delete(b, 1)
	fmt.Println(map_1, b)
	s := "anb"
	fmt.Println(s)
	s = ""
	fmt.Println(s, "hello")

}
