package main

import (
	"reflect"
	"testing"
)

//TestPascal is used to test Pascal program
func TestPascal(t *testing.T) {
	testcases := []struct {
		desc           string
		input          int
		expectedOutput [][]int
	}{
		{desc: "pass", input: 5, expectedOutput: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{desc: "pass", input: 4, expectedOutput: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}}},
	}
	for _, v := range testcases {
		output := Pascal(v.input)
		if !reflect.DeepEqual(output, v.expectedOutput) {
			t.Errorf(" actual output is %v but got %v ", output, v.expectedOutput)

		}
	}
}

//BenchmarkPascal testing is useds to test the performance of Pascal
func BenchmarkPascal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pascal(5)

	}
}
