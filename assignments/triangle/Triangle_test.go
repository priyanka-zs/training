package triangle

import (
	"reflect"
	"testing"
)

//TestTriangle is used to test the Triangle func
func TestTriangle(t *testing.T) {
	testcases := []struct {
		des            string
		input          []int
		expectedOutput string
	}{
		{des: "pass", input: []int{5, 4, 3}, expectedOutput: "scaleneTriangle"},
		{des: "pass", input: []int{5, 5, 5}, expectedOutput: "EquilateralTriangle"},
		{des: "pass", input: []int{5, 5, 4}, expectedOutput: "IsoscelesTriangle"},
		{des: "pass", input: []int{5, 5, 15}, expectedOutput: "given lengths cannot form a triangle"},
		{des: "pass", input: []int{5, 4, -3}, expectedOutput: "length cannot be negative"},
	}

	for _, v := range testcases {

		output := Triangle(v.input)

		if !(reflect.DeepEqual(output, v.expectedOutput)) {
			t.Errorf("testcase failed ,expected output is %v but got %v", v.expectedOutput, output)
		}
		/*if k == v.input && !(assert.Equal(t, v.expectedOutput, output)) {
			t.Errorf("testcase failed ,expected output is %v but got %v", v.expectedOutput, output)
		}*/
	}
}

//BenchmarkTriangle is used to test the efficiency of Triangle code
func BenchmarkTriangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Triangle([]int{5, 5, 5})
		Triangle([]int{5, 5, 4})
		Triangle([]int{5, 4, 3})
	}
}
