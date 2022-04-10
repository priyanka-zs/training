package prime

import (
	"reflect"
	"testing"
)

/*TestPrime testing is done on generating prime in a given range*/
func TestPrime(t *testing.T) {
	testcases := []struct {
		des            string
		input          int
		expectedOutput []int
	}{
		{des: "pass", input: 5, expectedOutput: []int{2, 3}},
		//{des: "pass", input: 20, expectedOutput: []int{2, 3, 5, 7, 11, 13, 17, 19}},
	}

	for _, v := range testcases {
		k := 5
		output := Prime(k)

		if !(reflect.DeepEqual(output, v.expectedOutput)) {
			t.Errorf("testcase failed ,expected output is %v but got %v", v.expectedOutput, output)
		}
		/*if k == v.input && !(assert.Equal(t, v.expectedOutput, output)) {
			t.Errorf("testcase failed ,expected output is %v but got %v", v.expectedOutput, output)
		}*/
	}
}

/*BenchmarkPrime is used to test the efficiency of code*/
func BenchmarkPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Prime(10)
		Prime(20)
	}
}
