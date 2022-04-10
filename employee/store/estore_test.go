package store

import (
	"context"
	"employee/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDB_CreateStore(t *testing.T) {
	emp := models.Employee{Id: 1, Name: "priya", Age: 20}
	testcases := []struct {
		desc   string
		input  models.Employee
		output models.Employee
		err    error
	}{
		{"", emp, emp, nil},
	}
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	s := New(db)
	ctx := context.Background()
	for _, tc := range testcases {
		mock.ExpectExec("insert into employee(id,name)VALUES (?,?)").
			WithArgs(emp.Id, emp.Name).WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(tc.err)
		_, err := s.CreateStore(ctx, tc.input)
		assert.Equal(t, tc.err, err)
	}
}
