package main

import (
	"fmt"
	"strconv"
)

func prac1() {
	a := "as12dsf34ff6"
	sum := 0
	s := ""
	for i, e := range a {
		if i == len(a)-1 {
			s = s + string(e)
			k, _ := strconv.Atoi(s)
			sum += k
		} else if !((string(e) >= string('a')) && (string(e) <= string('z'))) {
			s = s + string(e)
		} else {
			k, _ := strconv.Atoi(s)
			sum += k
			s = ""
		}
	}
	fmt.Println(sum)
}
