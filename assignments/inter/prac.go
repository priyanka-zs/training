package main

import (
	"fmt"
	"strconv"
)

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("my name is %v and i am( %v) years old", p.name, p.age)
}
func main() {
	p := person{"priyanka", 22}
	fmt.Println(p)
	i, err := strconv.Atoi("42")
	fmt.Println(i, err)
}
