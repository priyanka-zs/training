package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("num cannot be negative")
	}
	return math.Sqrt(x), nil
}
func main() {
	x, err := Sqrt(-10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}

}
