package geometry

import (
	"reflect"
	"testing"
)

//TestRectPeri is used to test the perimeter of a rectangle
func TestRectPeri(t *testing.T) {
	testcases := []struct {
		des            string
		input          rectangle
		expectedOutput float64
	}{
		{"to find perimeter of rect ", rectangle{10, 20}, 60.0},
		{"to find perimeter of rect ", rectangle{-2, 20}, 0.0},
	}
	for _, v := range testcases {

		output := (v.input).perimeter()

		if !(reflect.DeepEqual(output, v.expectedOutput)) {
			t.Errorf("testcase failed ,expected output is %v but got %v", v.expectedOutput, output)
		}
	}
}

//TestCirPeri is used to test the perimeter of a circle
func TestCirPeri(t *testing.T) {
	testcases := []struct {
		des            string
		input          circle
		expectedOutput float64
	}{
		{"to find perimeter of circle", circle(4), 25.12},
		{"to find perimeter of circle", circle(-4), 0.0},
	}
	for _, v := range testcases {

		output := (v.input).perimeter()

		if !(reflect.DeepEqual(output, v.expectedOutput)) {
			t.Errorf("testcase failed ,expected output is %v but got %v", v.expectedOutput, output)
		}
	}
}

//TestSqPeri is used to test the perimeter of a square
func TestSqPeri(t *testing.T) {
	testcases := []struct {
		des            string
		input          square
		expectedOutput float64
	}{
		{"to find perimeter of square", square(10), 40},
		{"to find perimeter of square", square(-10), 0.0},
	}
	for _, v := range testcases {

		output := (v.input).perimeter()

		if !(reflect.DeepEqual(output, v.expectedOutput)) {
			t.Errorf("testcase failed ,expected output is %v but got %v", v.expectedOutput, output)
		}
	}
}

//TestRectArea is used to best the area of a rectangle
func TestRectArea(t *testing.T) {
	testcases := []struct {
		des            string
		input          rectangle
		expectedOutput float64
	}{
		{"to find area of rectangle", rectangle{10, 20}, 200},
		{"to find area of rectangle", rectangle{-10, 20}, 0.0},
	}
	for _, v := range testcases {

		output := (v.input).area()

		if !(reflect.DeepEqual(output, v.expectedOutput)) {
			t.Errorf("testcase failed ,expected output is %v but got %v", v.expectedOutput, output)
		}
	}
}

//TestCircleArea is used to best the area of a circle
func TestCircleArea(t *testing.T) {
	testcases := []struct {
		des            string
		input          circle
		expectedOutput float64
	}{
		{"to find area of circle", circle(5), 78.5},
		{"to find area of circle", circle(-5), 0.0},
	}
	for _, v := range testcases {

		output := (v.input).area()

		if !(reflect.DeepEqual(output, v.expectedOutput)) {
			t.Errorf("testcase failed ,expected output is %v but got %v", v.expectedOutput, output)
		}
	}
}

//TestSqArea is used to best the area of a square
func TestSqArea(t *testing.T) {
	testcases := []struct {
		des            string
		input          square
		expectedOutput float64
	}{
		{"to find area of square", square(10), 100},
		{"to find area of square", square(-10), 0.0},
	}
	for _, v := range testcases {

		output := (v.input).area()

		if !(reflect.DeepEqual(output, v.expectedOutput)) {
			t.Errorf("testcase failed ,expected output is %v but got %v", v.expectedOutput, output)
		}
	}
}

//BenchmarkRectPeri is used to test the efficiency of code
func BenchmarkRectPeri(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rectangle{4, 5}.perimeter()
	}
}

//BenchmarkCirPeri is used to test the efficiency of code
func BenchmarkCirPeri(b *testing.B) {
	for i := 0; i < b.N; i++ {
		circle(5).perimeter()
	}
}

//BenchmarkSqPeri is used to test the efficiency of code
func BenchmarkSqPeri(b *testing.B) {
	for i := 0; i < b.N; i++ {
		square(5).perimeter()
	}
}

//BenchmarkRectArea is used to test the efficiency of code
func BenchmarkRectArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rectangle{4, 5}.area()
	}
}

//BenchmarkCirArea is used to test the efficiency of code
func BenchmarkCirArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		circle(5).area()
	}
}

//BenchmarkSqArea is used to test the efficiency of code
func BenchmarkSqArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		square(5).area()
	}
}
