package handler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Post(t *testing.T) {
	testcases := []struct {
		desc  string
		input string
		err   error
	}{
		{"success case", "1", nil},
	}
	for _, tc := range testcases {

		err := Post(tc.input)
		assert.Equal(t, tc.err, err)

	}

}
