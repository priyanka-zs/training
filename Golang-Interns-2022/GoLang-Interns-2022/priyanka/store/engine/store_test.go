package engine

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/zopsmart/GoLang-Interns-2022/model"
)

// TestDB_Create is used to test the create method of engine
func TestDB_Create(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	id1 := uuid.New()

	engine := model.Engine{EngineID: id1, Displacement: 12, NoOfCylinders: 5, Ranges: 0}

	testcases := []struct {
		desc     string
		input    *model.Engine
		exOutput *model.Engine
		err      error
	}{
		{"success case", &engine, &engine, nil},
		{"query error", &model.Engine{Displacement: 12, NoOfCylinders: 5, Ranges: 0}, &model.Engine{}, errors.New("error in insertion")},
	}

	mock.ExpectExec("INSERT INTO engine(Engine_Id,Displacement,No_of_cylinders,`Range`) "+
		"VALUES(?,?,?,?)").WithArgs(engine.EngineID, engine.Displacement, engine.NoOfCylinders, engine.Ranges).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO engine(Engine_Id,Displacement,No_of_cylinders,`Range`) "+
		"VALUES(?,?,?,?)").WithArgs(engine.EngineID, engine.Displacement, engine.NoOfCylinders, engine.Ranges).
		WillReturnError(errors.New("error in insertion"))

	s := New(db)

	for _, tc := range testcases {
		resp, err := s.EngineCreate(ctx, &engine)
		assert.Equal(t, tc.exOutput, resp)
		assert.Equal(t, err, tc.err)
	}
}

// TestDB_Get is used to test the get method of engine
func TestDB_Get(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	id1 := uuid.New()
	id2 := uuid.New()

	testcases := []struct {
		desc     string
		id       uuid.UUID
		exOutput *model.Engine
		err      error
	}{
		{"success case", id1, &model.Engine{EngineID: id1, Displacement: 20, NoOfCylinders: 3, Ranges: 0}, nil},
		{"invalid id", id2, &model.Engine{}, errors.New("query error")},
	}

	rows := sqlmock.NewRows([]string{"id", "displacement", "noOfCylinders", "ranges"})
	rows.AddRow(id1, 20, 3, 0)

	mock.ExpectQuery("SELECT Engine_Id,Displacement,No_of_cylinders,`Range` FROM engine").WillReturnRows(rows)
	mock.ExpectQuery("SELECT Engine_Id,Displacement,No_of_cylinders,`Range` FROM engine").WillReturnError(errors.New("query error"))

	s := New(db)

	for _, tc := range testcases {
		resp, err := s.EngineGet(ctx, tc.id)
		assert.Equal(t, err, tc.err)
		assert.Equal(t, tc.exOutput, resp)
	}
}

// TestDB_Update is used to test the delete method of engine
func TestDB_Update(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	id1 := uuid.New()
	engine := model.Engine{EngineID: id1, Displacement: 20, NoOfCylinders: 3, Ranges: 0}
	id2 := uuid.New()
	engine1 := model.Engine{EngineID: id2, Displacement: 20, NoOfCylinders: 3, Ranges: 0}

	testcases := []struct {
		desc     string
		input    *model.Engine
		exOutput *model.Engine
		err      error
	}{
		{"success case", &engine, &engine, nil},
		{"id not in db", &engine1, &model.Engine{}, errors.New("query error")},
	}

	mock.ExpectExec("UPDATE engine SET Displacement=?,No_of_cylinders=?,`Range`=? where Engine_Id=?").
		WithArgs(engine.Displacement, engine.NoOfCylinders, engine.Ranges, engine.EngineID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("UPDATE engine SET Displacement=?,No_of_cylinders=?,`Range`=? where Engine_Id=?").
		WithArgs(engine1.Displacement, engine1.NoOfCylinders, engine1.Ranges, engine1.EngineID).
		WillReturnError(errors.New("query error"))

	s := New(db)
	for _, tc := range testcases {
		resp, err := s.EngineUpdate(ctx, tc.input)
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.exOutput, resp)
	}
}

// TestDB_Delete is used to test the delete method of engine
func TestDB_Delete(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	id1 := uuid.New()
	testcases := []struct {
		desc string
		id   uuid.UUID
		err  error
	}{
		{"success case", id1, nil},
		{"id not in db", uuid.Nil, errors.New("query error")},
	}

	mock.ExpectExec("delete from engine where Engine_Id=?").WithArgs(id1.String()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("delete from engine where Engine_Id=?").WithArgs(uuid.Nil).
		WillReturnError(errors.New("query error"))

	s := New(db)
	for _, tc := range testcases {
		err := s.EngineDelete(ctx, tc.id)
		assert.Equal(t, tc.err, err)
	}
}
