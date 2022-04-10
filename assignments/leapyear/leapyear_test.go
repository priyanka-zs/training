package leapyear

import (
	"reflect"
	"testing"
)

//TestLeapYear is used to test the LeapYear func
func TestLeapYear(t *testing.T) {
	testcases := []struct {
		des            string
		input          int
		expectedOutput string
	}{
		{des: "pass", input: 2000, expectedOutput: "LeapYear"},
		{des: "pass", input: 1900, expectedOutput: "NotaLeapYear"},
		//{des: "pass", input: 20, expectedOutput: []int{2, 3, 5, 7, 11, 13, 17, 19}},
	}

	for _, v := range testcases {

		output := LeapYear(v.input)

		if !(reflect.DeepEqual(output, v.expectedOutput)) {
			t.Errorf("testcase failed ,expected output is %v but got %v", v.expectedOutput, output)
		}
	}
}
