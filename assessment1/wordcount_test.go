package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_WordCount(t *testing.T) {

	testcases := []struct {
		desc   string
		input  string
		output int
	}{
		{"success case", " I went to the beach with my dog yesterday. My dog had a good time.", 2},
	}
	for _, tc := range testcases {
		res := WordCount(tc.input)
		assert.Equal(t, tc.output, res)
	}

}
