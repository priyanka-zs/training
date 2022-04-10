package main

import (
	"fmt"
	"math"
)

/*type vertex struct {
	x, y float64
}

func (v vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func main() {
	v := vertex{3, 4}
	fmt.Println(v.abs())

}*/

/*type MyInt int

func (i MyInt) f1() int {
	if i < 0 {
		return int(-i)
	}
	return int(i)
}
func main(){
	v := MyInt(-10)
	fmt.Println(v.f1())
}*/
type vertex struct {
	x, y float64
}

func (v vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func (v *vertex) f1(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}
func main() {
	v := vertex{3, 4}
	fmt.Println(v)
	v.f1(10)
	fmt.Println(v)
	fmt.Println(v.abs())

}
