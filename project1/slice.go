package main

import "fmt"

func slice() {
	/*b := []byte{'a', 'e', 'i', 'o', 'u'}
	fmt.Println(string(b[0:2]))
	s := "acd@rg#%&*"
	res := ""
	res1 := ""
	for _, e := range s {
		if !((string(e) >= string('a')) && (string(e) <= string('z'))) {
			res += string(e)
		} else {
			res1 += string(e)
		}
	}
	fmt.Println(res + res1)*/
	a := [5]int{1, 2, 3, 4, 5}
	s := a[0:3]
	s[0] = 12
	fmt.Println(a)

}
