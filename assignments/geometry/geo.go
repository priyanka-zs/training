package geometry

//shape interface has perimeter and area methods
type shape interface {
	perimeter() float64
	area() float64
}

//rectangle has length and breadth variables in it
type rectangle struct {
	l float64
	b float64
}

type circle float64
type square float64

//perimeter takes rectangle as receiver and returns the perimeter of rectangle
func (r rectangle) perimeter() float64 {
	if r.l <= 0 || r.b <= 0 {
		return 0.0
	}
	return 2 * (r.l + r.b)
}

//perimeter takes Square as receiver and returns the perimeter of square
func (s square) perimeter() float64 {
	if s <= 0 {
		return 0.0
	}
	return float64(4 * s)
}

//perimeter takes Circle as receiver and  returns the perimeter of circle
func (c circle) perimeter() float64 {
	if c <= 0 {
		return 0.0
	}
	return float64(2 * 3.14 * c)
}

//area takes Square as receiver and returns the area of square
func (s square) area() float64 {
	if s <= 0 {
		return 0.0
	}
	return float64(s * s)
}

//area takes Rectangle as receiver and returns the area of rectangle
func (r rectangle) area() float64 {
	if r.l <= 0 || r.b <= 0 {
		return 0.0
	}
	return r.l * r.b
}

//area takes circle as receiver and returns the area of circle
func (c circle) area() float64 {
	if c <= 0 {
		return 0.0
	}
	return float64(3.14 * c * c)
}
