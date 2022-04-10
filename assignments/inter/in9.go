package main

import "fmt"

type Abser interface {
	abs() float64
}
type myFloat float64
type Vertex struct {
	x float64
	y float64
}

func (f myFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)

}
func (v Vertex) abs() float64 {
	return v.x + v.y
}
func main() {
	var a Abser
	v := Vertex{2.5, 3.5}
	a = v
	fmt.Println(a.abs())
	f := myFloat(-2.5)
	a = f
	fmt.Println(a.abs())
}