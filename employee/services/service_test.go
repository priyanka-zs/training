package services

import (
	"employee/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
)

func Test_Get(t *testing.T) {
	emp := models.Employee{Id: 1, Name: "priya", Age: 27}
	testcases := []struct {
		desc   string
		input  int
		output bool
	}{
		{"success", 27, true},
	}
	ctrl := gomock.NewController(t)
	MockEmp := NewMockEmp(ctrl)
	s := New(MockEmp)
	ctx := context.Background()
	for _, tc := range testcases {
		MockEmp.EXPECT().GetById(gomock.Any(), tc.input).Return(emp, nil)
		res := s.Get(ctx, tc.input)
		assert.Equal(t, tc.output, res)
	}
}
