package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abc@r#d"
	s2 := ""
	for i := len(s) - 1; i >= 0; i-- {

		if string(s[i]) != "@" && string(s[i]) != "#" && string(s[i]) != "$" &&
			string(s[i]) != "%" && string(s[i]) != "^" && string(s[i]) != "&" && string(s[i]) != "*" && string(s[i]) != "?" {

		}
	}
	fmt.Println(s2)

	for i := 0; i < len(s); i++ {

		if string(s[i]) == "@" || string(s[i]) == "#" || string(s[i]) == "$" ||
			string(s[i]) == "%" && string(s[i]) == "^" || string(s[i]) == "?" {
			s2 = strings.Replace(s2, string(s2[i-1]), string(s2[i-1])+string(s[i]), 1)
		}
	}
	//	fmt.Println(s2)
}
