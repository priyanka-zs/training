package bubblesort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BubbleSort(t *testing.T) {
	testcases := []struct {
		desc   string
		input  []int
		output []int
	}{
		{"success", []int{1, 2, 3}, []int{1, 2, 3}},
	}
	for _, tc := range testcases {
		arr := BubbleSort(tc.input)
		assert.Equal(t, tc.output, arr)
	}
}
func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort([]int{1, 2, 3})
	}
}
