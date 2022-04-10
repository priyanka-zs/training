package engine

import (
	"context"
	_ "database/sql"
	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gofr_cardealership/model"
	"log"
	"testing"
)

func NewMock() (*gofr.Context, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	ctx := gofr.NewContext(nil, nil, &gofr.Gofr{DataStore: datastore.DataStore{ORM: db}})
	ctx.Context = context.Background()

	return ctx, mock
}

// TestEngineCreate is used to test the create method of engine
func TestEngineCreate(t *testing.T) {
	ctx, mock := NewMock()

	id1 := uuid.New()

	engine := model.Engine{EngineID: id1, Displacement: 12, NoOfCylinders: 5, Ranges: 0}

	testcases := []struct {
		desc     string
		input    *model.Engine
		exOutput *model.Engine
		err      error
	}{
		{"success case", &engine, &engine, nil},
		{"query error", &model.Engine{Displacement: 12, NoOfCylinders: 5, Ranges: 0}, &model.Engine{}, errors.Error("Internal server error")},
	}

	mock.ExpectExec("INSERT INTO Engine(EngineID,Displacement,NoOfCylinders,Ranges) "+
		"VALUES(?,?,?,?)").WithArgs(engine.EngineID, engine.Displacement, engine.NoOfCylinders, engine.Ranges).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO Engine(EngineID,Displacement,NoOfCylinders,Ranges) "+
		"VALUES(?,?,?,?)").WithArgs(engine.EngineID, engine.Displacement, engine.NoOfCylinders, engine.Ranges).
		WillReturnError(errors.Error("Internal server error"))

	s := New()

	for _, tc := range testcases {
		_, err := s.EngineCreate(ctx, &engine)
		assert.Equal(t, tc.err, err)
	}
}

// TestDB_Get is used to test the get method of engine
func TestDB_Get(t *testing.T) {
	ctx, mock := NewMock()

	id1 := uuid.MustParse("045b658e-9160-4f55-8e5a-be8ceb13fbf5")
	id2 := uuid.New()

	testcases := []struct {
		desc     string
		id       uuid.UUID
		exOutput *model.Engine
		err      error
	}{
		{"success case", id1, &model.Engine{EngineID: id1, Displacement: 20, NoOfCylinders: 3, Ranges: 0}, nil},
		{"invalid id", id2, &model.Engine{}, errors.Error("Internal server error")},
	}

	rows := sqlmock.NewRows([]string{"EngineID", "Displacement", "NoOfCylinders", "Ranges"})
	rows.AddRow(id1, 20, 3, 0)

	mock.ExpectQuery("SELECT EngineID,Displacement,NoOfCylinders,Ranges FROM Engine where EngineID = ?").WithArgs(id1).WillReturnRows(rows)
	mock.ExpectQuery("SELECT EngineID,Displacement,NoOfCylinders,Ranges FROM Engine where EngineID = ?").WithArgs(id1).WillReturnError(errors.Error("query error"))

	s := New()

	for _, tc := range testcases {
		_, err := s.EngineGet(ctx, tc.id)
		assert.Equal(t, tc.err, err)
	}
}

// TestDB_Delete is used to test the delete method of engine
func TestDB_Delete(t *testing.T) {
	ctx, mock := NewMock()

	id1 := uuid.New()
	testcases := []struct {
		desc string
		id   uuid.UUID
		err  error
	}{
		{"success case", id1, nil},
		{"id not in db", uuid.Nil, errors.Error("query error")},
	}

	mock.ExpectExec("delete from Engine where EngineID=?").WithArgs(id1.String()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("delete from Engine where EngineID=?").WithArgs(uuid.Nil).
		WillReturnError(errors.Error("query error"))

	s := New()
	for _, tc := range testcases {
		err := s.EngineDelete(ctx, tc.id)
		assert.Equal(t, tc.err, err)
	}
}

// TestDB_Update is used to test the delete method of engine
func TestDB_Update(t *testing.T) {
	ctx, mock := NewMock()

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
		{"id not in db", &engine1, &model.Engine{}, errors.Error("query error")},
	}

	mock.ExpectExec("UPDATE Engine SET Displacement=?,NoOfCylinders=?,Ranges=? where EngineID=?").
		WithArgs(engine.Displacement, engine.NoOfCylinders, engine.Ranges, engine.EngineID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("UPDATE Engine SET Displacement=?,NoOfCylinders=?,Ranges=? where EngineID=?").
		WithArgs(engine1.Displacement, engine1.NoOfCylinders, engine1.Ranges, engine1.EngineID).
		WillReturnError(errors.Error("query error"))

	s := New()
	for _, tc := range testcases {
		_, err := s.EngineUpdate(ctx, tc.input)
		assert.Equal(t, tc.err, err)
	}
}
