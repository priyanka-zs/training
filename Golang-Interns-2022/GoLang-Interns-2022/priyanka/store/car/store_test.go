package car

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/zopsmart/GoLang-Interns-2022/model"
)

// TestDB_CarCreate is used to test the CarCreate method
func TestDB_CarCreate(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error %s was not expected when opening a stub database connection", err)
	}

	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			return
		}
	}(db)

	car := model.Car{Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol"}

	testcases := []struct {
		desc  string
		input *model.Car
		err   error
	}{
		{"success case", &car, nil},
		{"query error", &car, errors.New("error in insertion")},
	}
	mock.ExpectExec("INSERT INTO car(Id,Name,Year,Brand,FuelType,Engine_Id)"+
		"VALUES(?,?,?,?,?,?)").WithArgs(sqlmock.AnyArg(), car.Name, car.Year, car.Brand, car.Fuel, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO car(Id,Name,Year,Brand,FuelType,Engine_Id)"+
		"VALUES(?,?,?,?,?,?)").WithArgs(sqlmock.AnyArg(), car.Name, car.Year, car.Brand, car.Fuel, sqlmock.AnyArg()).
		WillReturnError(errors.New("error in insertion"))

	s := New(db)
	for _, tc := range testcases {
		_, err := s.CarCreate(ctx, tc.input)
		assert.Equal(t, tc.err, err)
	}
}

// TestDB_Get is used to test the get method of car
func TestDB_Get(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	s := New(db)
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			return
		}
	}(db)

	id1, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("Id generation failed")
	}

	id2, _ := uuid.Parse("163744be-d162-441b-9386-43ab6042ec6d")

	testcases := []struct {
		desc     string
		id       uuid.UUID
		exOutput model.Car
		err      error
	}{
		{"success case", id1, model.Car{ID: id1, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol", Engine: model.Engine{}}, nil},
		{"invalid id", id2, model.Car{}, errors.New("error in get")},
	}

	rows := sqlmock.NewRows([]string{"Id", "Name", "Year", "Brand", "FuelType", "Engine_Id"})
	rows.AddRow(id1, "bmwZ4", 2000, "bmw", "petrol", uuid.Nil)
	mock.ExpectQuery("SELECT Id,Name,Year,Brand,FuelType,Engine_Id FROM car where Id=?").
		WithArgs(id1).WillReturnRows(rows)
	mock.ExpectQuery("SELECT Id,Name,Year,Brand,FuelType,Engine_Id FROM car where Id=?").
		WithArgs(id2).WillReturnError(errors.New("error in get"))

	for _, tc := range testcases {
		resp, err := s.CarGet(ctx, tc.id)
		assert.Equal(t, &tc.exOutput, resp)
		assert.Equal(t, tc.err, err)
	}
}

// TestDB_Update is used to test the update method of car
func TestDB_Update(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	s := New(db)

	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			return
		}
	}(db)

	id1 := uuid.New()
	id2, err := uuid.Parse("023744be-d162-441b-9386-43ab6002ec5d")

	if err != nil {
		return
	}

	car := model.Car{ID: id1, Name: "bmw", Year: 2000, Brand: "bmw", Fuel: "petrol", Engine: model.Engine{}}
	car1 := model.Car{ID: id2, Name: "ferrari f50", Year: 2000, Brand: "ferrari", Fuel: "petrol", Engine: model.Engine{}}

	testcases := []struct {
		desc     string
		input    *model.Car
		exOutput *model.Car
		err      error
	}{
		{"success case", &car, &car, nil},
		{"id not in db", &car1, &model.Car{}, errors.New("id not in db")},
	}

	mock.ExpectExec("UPDATE car SET Name=?,Year=?,Brand=?,FuelType=? where Id=?").
		WithArgs(car.Name, car.Year, car.Brand, car.Fuel, car.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("UPDATE car SET Name=?,Year=?,Brand=?,FuelType=? where Id=?").
		WithArgs(car1.Name, car1.Year, car1.Brand, car1.Fuel, car1.ID).
		WillReturnError(errors.New("id not in db"))

	for _, tc := range testcases {
		resp, err := s.CarUpdate(ctx, tc.input)
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.exOutput, resp)
	}
}

// TestDB_Delete is used to test the delete method of car
func TestDB_Delete(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	s := New(db)
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			return
		}
	}(db)

	id1 := uuid.New()
	id2 := uuid.New()

	testcases := []struct {
		desc string
		id   uuid.UUID
		err  error
	}{
		{"success case", id1, nil},
		{"id not in db", id2, errors.New("error in deletion")},
	}

	mock.ExpectExec("delete from car where Id=?").WithArgs(id1.String()).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("delete from car where Id=?").WithArgs(id2.String()).WillReturnError(errors.New("error in deletion"))

	for _, tc := range testcases {
		err = s.CarDelete(ctx, tc.id)
		assert.Equal(t, tc.err, err)
	}
}

// TestDB_GetByBrand is used to test the get method of car
func TestDB_GetByBrandSuccess(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	s := New(db)

	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			return
		}
	}(db)

	id1 := uuid.New()
	id2 := uuid.New()
	car := []*model.Car{{ID: id1, Name: "ferrari f40", Year: 2015, Brand: "ferrari", Fuel: "diesel", Engine: model.Engine{EngineID: id2}}}
	rows := sqlmock.NewRows([]string{"id", "name", "year", "brand", "fuelType", "engine_id"})
	rows.AddRow(id1, "ferrari f40", 2015, "ferrari", "diesel", id1)

	testcases := []struct {
		desc     string
		brand    string
		isEngine bool
		quErr    error
		exOutput []*model.Car
		err      error
	}{
		{"success case", "ferrari", false, nil, car, nil},
	}
	for _, tc := range testcases {
		mock.ExpectQuery("select * from car where Brand=?").WillReturnRows(rows)

		_, err = s.CarGetByBrand(context.TODO(), tc.brand, tc.isEngine)

		assert.Equal(t, tc.err, err)
	}
}
func TestDB_GetByBrandFailure(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	s := New(db)

	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			return
		}
	}(db)

	id1 := uuid.New()
	emptyRows := sqlmock.NewRows([]string{"Id", "Name", "Year", "Brand", "fuelType", "Engine_ID"})
	scanErrRow := sqlmock.NewRows([]string{"Id", "Name", "Year", "Fuel", "Engine_ID"}).AddRow(id1.String(),
		"bmw 2", 2018, "electric", uuid.New())
	testcases := []struct {
		desc       string
		brand      string
		isEngine   bool
		quErr      error
		outputRows *sqlmock.Rows
		exOutput   []*model.Car
		err        error
	}{
		{"failure", "ferrari", false, errors.New("query error"),
			emptyRows, []*model.Car{}, errors.New("query error")},
		{"scan error", "ferrari", false, nil,
			scanErrRow, []*model.Car{}, errors.New("error while scanning")},
	}

	for _, tc := range testcases {
		mock.ExpectQuery("select * from car where Brand=?").WithArgs(tc.brand).WillReturnRows(tc.outputRows).WillReturnError(tc.quErr)

		_, err = s.CarGetByBrand(context.TODO(), tc.brand, tc.isEngine)

		assert.Equal(t, tc.err, err)
	}
}
