package triangle

//Triangle func is used to check the type of triangle and return the type of triangle
func Triangle(a []int) string {
	if a[0]+a[1] < a[2] || a[1]+a[2] < a[0] || a[0]+a[2] < a[1] {
		return "given lengths cannot form a triangle"
	} else if a[0] <= 0 || a[1] <= 0 || a[2] <= 0 {
		return "length cannot be negative"
	} else if a[0] == a[1] && a[1] == a[2] {
		return "EquilateralTriangle"
	} else if a[0] == a[1] || a[1] == a[2] || a[2] == a[0] {
		return "IsoscelesTriangle"
	} else {
		return "scaleneTriangle"
	}
}
