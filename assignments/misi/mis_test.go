package main

import (
	"reflect"
	"testing"
)

func TestPrime(t *testing.T) {
	testcases := []struct {
		des            string
		input          string
		expectedOutput map[string]int
	}{
		{des: "pass", input: "Mississippi", expectedOutput: map[string]int{"M": 1, "i": 4, "p": 2, "s": 4}},
		//{des: "fail", input: "", expectedOutput: nil},

		//{des: "pass", input: 20, expectedOutput: []int{2, 3, 5, 7, 11, 13, 17, 19}},
	}

	for _, v := range testcases {

		output := mis(v.input)

		if !(reflect.DeepEqual(output, v.expectedOutput)) {
			t.Errorf("testcase failed ,expected output is %v but got %v", v.expectedOutput, output)
		}
		/*if k == v.input && !(assert.Equal(t, v.expectedOutput, output)) {
			t.Errorf("testcase failed ,expected output is %v but got %v", v.expectedOutput, output)

		}*/

	}
}
