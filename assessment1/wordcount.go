package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) int {
	return strings.Count(s, "dog")

}
func main() {
	s := " I went to the beach with my dog yesterday. My dog had a good time."
	fmt.Println(WordCount(s))
}
