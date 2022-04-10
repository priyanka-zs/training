package main

import "fmt"

type Shape interface {
	area() float64
}
type circle struct {
	r float64
}
type rectangle struct {
	l float64
	b float64
}

func (c *circle) area() float64 {
	return 3.14 * c.r * c.r
}
func (r rectangle) area() float64 {
	return r.l * r.b
}
func main() {
	c := &circle{5.2}
	r := &rectangle{5, 2}
	/*var s Shape= c
	var s1 Shape= r
	fmt.Println(s.area(), s1.area())*/
	shapes := []Shape{c, r}
	for _, shape := range shapes {
		fmt.Println(shape.area())

	}

}
