package stores

import (
	"context"
	"errors"
	"petstore/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Test_Post is used to test the post method
func Test_Post(t *testing.T) {
	id := uuid.New()
	pet := models.Pet{ID: id, Name: "bruno", Age: 2, Species: "dog", Status: "Available"}
	testcases := []struct {
		desc   string
		input  models.Pet
		output models.Pet
		err    error
	}{
		{"success", pet, pet, nil},
		{"error in store layer", models.Pet{}, models.Pet{}, errors.New("error in store layer")},
	}
	ctx := context.Background()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		return
	}

	s := New(db)

	for _, tc := range testcases {
		mock.ExpectExec("insert into pet(Id,Name,Age,Species,Status)VALUES (?,?,?,?,?)").
			WithArgs(sqlmock.AnyArg(), pet.Name, pet.Age, pet.Species, pet.Status).
			WillReturnResult(sqlmock.NewResult(1, 1))

		_, err := s.Post(ctx, tc.input)

		assert.Equal(t, tc.err, err)
	}
}

// Test_GetByID is used to test the GetByID in store layer
func Test_GetByID(t *testing.T) {
	id := uuid.New()
	pet := models.Pet{ID: id, Name: "bruno", Age: 2, Species: "dog", Status: "Available"}
	testcases := []struct {
		desc   string
		input  uuid.UUID
		output models.Pet
		err    error
	}{
		{"success", id, pet, nil},
		{"nil id", uuid.Nil, models.Pet{}, errors.New("error in store layer")},
	}
	ctx := context.Background()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		return
	}

	s := New(db)

	rows := sqlmock.NewRows([]string{"Id", "Name", "Age", "Species", "Status"})
	rows.AddRow(pet.ID, pet.Name, pet.Age, pet.Species, pet.Status)

	for _, tc := range testcases {
		mock.ExpectQuery("select * from pet where id=?").WithArgs(id).WillReturnRows(rows)

		_, err := s.GetByID(ctx, tc.input)

		if err != nil {
			return
		}

		assert.Equal(t, tc.err, err)
	}
}

func Test_Delete(t *testing.T) {
	id := uuid.New()
	testcases := []struct {
		desc  string
		input uuid.UUID
		err   error
	}{
		{"success case", id, nil},
	}
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	s := New(db)
	ctx := context.Background()

	for _, tc := range testcases {
		mock.ExpectExec("delete from pet where Id=?").WithArgs(tc.input).WillReturnResult(sqlmock.NewResult(1, 1))
		err := s.Delete(ctx, tc.input)
		assert.Equal(t, tc.err, err)
	}
}

func Test_DeleteErr(t *testing.T) {
	testcases := []struct {
		desc  string
		input uuid.UUID
		err   error
	}{
		{"nil id", uuid.Nil, errors.New("nil id")},
	}
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	s := New(db)
	ctx := context.Background()

	for _, tc := range testcases {
		mock.ExpectExec("delete from pet where Id=?").WithArgs(tc.input).WillReturnError(tc.err)
		err := s.Delete(ctx, tc.input)
		assert.Equal(t, tc.err, err)
	}
}
