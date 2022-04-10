package services

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"

	"github.com/zopsmart/GoLang-Interns-2022/model"
)

// TestValidation is used to test the validate function
func TestValidation(t *testing.T) {
	id := uuid.New()
	testcases := []struct {
		desc  string
		input model.Car
		err   error
	}{
		{desc: "success case", input: model.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol",
			Engine: model.Engine{EngineID: id, Displacement: 20, NoOfCylinders: 5, Ranges: 0}},
			err: nil},
		{desc: "year not in the given limit",
			input: model.Car{ID: id, Name: "bmwZ4", Year: 1970, Brand: "bmw", Fuel: "petrol",
				Engine: model.Engine{EngineID: id, Displacement: 20, NoOfCylinders: 5, Ranges: 0}},
			err: errors.New("year should be between 1980 and 2022")},
		{desc: "invalid brand", input: model.Car{ID: id, Name: "tataZ4", Year: 2000, Brand: "tata", Fuel: "petrol",
			Engine: model.Engine{EngineID: id, Displacement: 20, NoOfCylinders: 5, Ranges: 0}},
			err: errors.New("invalid brand")},
		{desc: "invalid fuel type", input: model.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "oil",
			Engine: model.Engine{EngineID: id, Displacement: 20, NoOfCylinders: 5, Ranges: 0}},
			err: errors.New("invalid fuel")},
		{desc: "displacement cannot be negative", input: model.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol",
			Engine: model.Engine{EngineID: id, Displacement: -2, NoOfCylinders: 5, Ranges: 0}},
			err: errors.New("displacement should be positive")},
		{desc: "noOfCylinders cannot be negative", input: model.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol",
			Engine: model.Engine{EngineID: id, Displacement: 2, NoOfCylinders: -5, Ranges: 0}},
			err: errors.New("noOfCylinders should be positive")},
		{desc: "range cannot be negative", input: model.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "electric",
			Engine: model.Engine{EngineID: id, Displacement: 0, NoOfCylinders: 0, Ranges: -1}},
			err: errors.New("range should be positive")},
		{desc: "range cannot be for petrol type", input: model.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol",
			Engine: model.Engine{EngineID: id, Displacement: 10, NoOfCylinders: 2, Ranges: 3}},
			err: errors.New("given fuel type does not support range")},
		{desc: "electric cannot have displacement", input: model.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "electric",
			Engine: model.Engine{EngineID: id, Displacement: 10, NoOfCylinders: 0, Ranges: 1}},
			err: errors.New("electric cannot have displacement")},
		{desc: "electric cannot have  noOfCylinders", input: model.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "electric",
			Engine: model.Engine{EngineID: id, Displacement: 0, NoOfCylinders: 10, Ranges: 1}},
			err: errors.New("electric cannot have  noOfCylinders")},
	}

	for _, tc := range testcases {
		err := Validation(tc.input)
		assert.Equal(t, err, tc.err)
	}
}

// TestCreateValidation is used to validate the input of create method
func TestCreateValidation(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	id1 := uuid.New()
	testcases := []struct {
		desc  string
		input model.Car
		err   error
	}{
		{desc: "invalid brand", input: model.Car{ID: id, Name: "tataZ4", Year: 2000, Brand: "tata", Fuel: "petrol",
			Engine: model.Engine{EngineID: id1, Displacement: 20, NoOfCylinders: 5, Ranges: 0}},
			err: errors.New("invalid brand")},
	}
	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		_, err := svc.Create(ctx, &tc.input)
		assert.Equal(t, err, tc.err)
	}
}

// TestCreateSuccess is used to test the Success case of Create
func TestCreateSuccess(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	id1 := uuid.New()
	testcases := []struct {
		desc     string
		input    model.Car
		carError error
		enError  error
		err      error
	}{
		{desc: "success case", input: model.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol",
			Engine: model.Engine{EngineID: id1, Displacement: 20, NoOfCylinders: 5, Ranges: 0}}, carError: nil, enError: nil, err: nil},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		mockCar.EXPECT().CarCreate(ctx, &tc.input).Return(&tc.input, tc.carError)
		mockEngine.EXPECT().EngineCreate(ctx, &tc.input.Engine).Return(&tc.input.Engine, tc.enError)
		_, err := svc.Create(ctx, &tc.input)
		assert.Equal(t, err, tc.err)
	}
}

// TestCreateCarError is used to test the Create method with error in store layer
func TestCreateCarError(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	id1 := uuid.New()
	testcases := []struct {
		desc     string
		input    model.Car
		carError error
		enError  error
		err      error
	}{

		{desc: "error in car store layer", input: model.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol",
			Engine: model.Engine{EngineID: id1, Displacement: 20, NoOfCylinders: 5, Ranges: 0}},
			carError: errors.New("error in car layer"), enError: nil, err: errors.New("error in car layer")},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		mockCar.EXPECT().CarCreate(ctx, &tc.input).Return(&tc.input, tc.carError)
		_, err := svc.Create(ctx, &tc.input)
		assert.Equal(t, err, tc.err)
	}
}

// TestCreateEngineError is used to test the create method with error in store layer
func TestCreateEngineError(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	id1 := uuid.New()
	testcases := []struct {
		desc     string
		input    model.Car
		carError error
		enError  error
		err      error
	}{

		{desc: "error in engine store layer", input: model.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol",
			Engine: model.Engine{EngineID: id1, Displacement: 20, NoOfCylinders: 5, Ranges: 0}},
			carError: nil, enError: errors.New("error in engine layer"), err: errors.New("error in engine layer")},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		mockCar.EXPECT().CarCreate(ctx, &tc.input).Return(&tc.input, tc.carError)
		mockEngine.EXPECT().EngineCreate(ctx, &tc.input.Engine).Return(&tc.input.Engine, tc.enError)
		_, err := svc.Create(ctx, &tc.input)
		assert.Equal(t, err, tc.err)
	}
}

// TestGetByIdSuccess is used to test the success case of GetByID
func TestGetByIdSuccess(t *testing.T) {
	ctx := context.Background()
	id1, _ := uuid.Parse("063744be-d162-441b-9386-43ab6042ed6f")
	id2, _ := uuid.Parse("063744be-d162-441b-9386-43ab6042ec5d")

	testcases := []struct {
		desc     string
		id       uuid.UUID
		exOutput model.Car
		err      error
	}{
		{"success case", id1, model.Car{ID: id1, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol",
			Engine: model.Engine{EngineID: id2, Displacement: 20, NoOfCylinders: 5}}, nil},
	}
	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		mockEngine.EXPECT().EngineGet(ctx, tc.exOutput.Engine.EngineID).Return(&tc.exOutput.Engine, tc.err)
		mockCar.EXPECT().CarGet(ctx, tc.id).Return(&tc.exOutput, tc.err)
		resp, _ := svc.GetByID(ctx, tc.id)
		assert.Equal(t, resp, &tc.exOutput)
	}
}

// TestGetByID is used to check getById method
func TestGetByNilID(t *testing.T) {
	ctx := context.Background()

	testcases := []struct {
		desc     string
		id       uuid.UUID
		exOutput model.Car
		err      error
	}{

		{"invalid id", uuid.Nil, model.Car{}, errors.New("missing id")},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		mockCar.EXPECT().CarGet(ctx, tc.id).Return(&tc.exOutput, tc.err)
		resp, _ := svc.GetByID(ctx, tc.id)
		assert.Equal(t, resp, &tc.exOutput)
	}
}

// TestGetByID is used to check getById method
func TestGetByID(t *testing.T) {
	ctx := context.Background()
	id, _ := uuid.NewRandom()

	testcases := []struct {
		desc     string
		id       uuid.UUID
		exOutput model.Car
		err1     error
		err2     error
	}{
		{"engine id doesn't exist", id, model.Car{}, nil, errors.New("invalid engine id")},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		mockCar.EXPECT().CarGet(ctx, tc.id).Return(&tc.exOutput, tc.err1)
		mockEngine.EXPECT().EngineGet(ctx, uuid.Nil).Return(&tc.exOutput.Engine, tc.err2)
		resp, _ := svc.GetByID(ctx, tc.id)
		assert.Equal(t, resp, &tc.exOutput)
	}
}

// TestGetSuccess is used to test the Get success case
func TestGetSuccess(t *testing.T) { //nolint
	ctx := context.Background()
	id1, _ := uuid.Parse("063744be-d162-441b-9386-43ab6042ec5d")

	car2 := []*model.Car{{ID: id1, Name: "ferrari f40", Year: 2015, Brand: "ferrari", Fuel: "diesel",
		Engine: model.Engine{EngineID: id1, Displacement: 20, NoOfCylinders: 2, Ranges: 0}}}

	testcases := []struct {
		desc     string
		id       uuid.UUID
		brand    string
		engine   bool
		exOutput []*model.Car
		err      error
	}{
		{"success case", id1, "ferrari", true, car2, nil},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for i, tc := range testcases {
		mockCar.EXPECT().CarGetByBrand(ctx, tc.brand, tc.engine).Return(tc.exOutput, tc.err)
		mockEngine.EXPECT().EngineGet(ctx, tc.id).Return(&tc.exOutput[i].Engine, tc.err)
		resp, err := svc.Get(ctx, tc.brand, tc.engine)
		assert.Equal(t, &resp, &tc.exOutput)
		assert.Equal(t, err, tc.err)
	}
}

// TestGet is used to test Get method
func TestGet(t *testing.T) { //nolint
	ctx := context.Background()
	id1, _ := uuid.Parse("063744be-d162-441b-9386-43ab6042ec5d")

	testcases := []struct {
		desc     string
		id       uuid.UUID
		brand    string
		engine   bool
		exOutput []*model.Car
		err      error
	}{
		{"error in store layer", id1, "ferrari", false, []*model.Car{}, errors.New("error from database layer")},
		{"invalid brand", id1, "tata", false, []*model.Car{}, errors.New("invalid brand")},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for i, tc := range testcases {
		if tc.brand == "bmw" || tc.brand == "tesla" || tc.brand == "ferrari" || tc.brand == "porsche" || tc.brand == "mercedes" {
			mockCar.EXPECT().CarGetByBrand(ctx, tc.brand, tc.engine).Return(tc.exOutput, tc.err)

			if tc.engine == true {
				mockEngine.EXPECT().EngineGet(ctx, tc.id).Return(&tc.exOutput[i].Engine, tc.err)
			}
		}

		resp, err := svc.Get(ctx, tc.brand, tc.engine)
		assert.Equal(t, &resp, &tc.exOutput)
		assert.Equal(t, err, tc.err)
	}
}

// TestGetEngineError is used to test the GetByBrand with error in engine get
func TestGetEngineError(t *testing.T) { //nolint
	ctx := context.Background()
	id1 := uuid.New()

	car2 := []*model.Car{{ID: id1, Name: "ferrari f40", Year: 2015, Brand: "ferrari", Fuel: "diesel",
		Engine: model.Engine{EngineID: id1, Displacement: 20, NoOfCylinders: 2, Ranges: 0}}}

	testcases := []struct {
		desc     string
		id       uuid.UUID
		brand    string
		engine   bool
		exOutput []*model.Car
		err      error
	}{

		{"error in engine get", id1, "tesla", true, car2, errors.New("error in engine get")},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		mockCar.EXPECT().CarGetByBrand(ctx, tc.brand, tc.engine).Return(tc.exOutput, nil)
		mockEngine.EXPECT().EngineGet(ctx, tc.id).Return(nil, tc.err)
		_, err := svc.Get(ctx, tc.brand, tc.engine)
		assert.Equal(t, err, tc.err)
	}
}

// TestUpdateValidation is used to test the update method
func TestUpdateValidation(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()

	testcases := []struct {
		desc     string
		input    model.Car
		exOutput model.Car
		err      error
	}{

		{desc: "year not in the given limit", input: model.Car{ID: id, Name: "bmwZ4", Year: 1970, Brand: "bmw", Fuel: "petrol",
			Engine: model.Engine{EngineID: id, Displacement: 20, NoOfCylinders: 5, Ranges: 0}},
			exOutput: model.Car{}, err: errors.New("year should be between 1980 and 2022")},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		resp, err := svc.Update(ctx, &tc.input)
		assert.Equal(t, resp, &tc.exOutput)
		assert.Equal(t, err, tc.err)
	}
}

// TestUpdate is used to test Update method
func TestUpdate(t *testing.T) {
	ctx := context.Background()
	id1, _ := uuid.Parse("063744be-d162-441b-9386-43ab6042ec5d")
	car := model.Car{ID: id1, Name: "bmw", Year: 2000, Brand: "bmw", Fuel: "petrol",
		Engine: model.Engine{EngineID: id1, Displacement: 20, NoOfCylinders: 5, Ranges: 0}}

	testcases := []struct {
		desc     string
		input    model.Car
		exOutput model.Car

		err error
	}{
		{"success case", car, car, nil},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		mockCar.EXPECT().CarGet(ctx, tc.input.ID).Return(&tc.exOutput, tc.err)
		mockEngine.EXPECT().EngineUpdate(ctx, &tc.input.Engine).Return(&tc.exOutput.Engine, tc.err)
		mockCar.EXPECT().CarUpdate(ctx, &tc.input).Return(&tc.exOutput, tc.err)

		resp, err := svc.Update(ctx, &tc.input)
		assert.Equal(t, resp, &tc.exOutput)
		assert.Equal(t, err, tc.err)
	}
}

// TestCarUpdateError is used to test the Update method in service layer
func TestCarUpdateError(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	id1 := uuid.New()
	car := model.Car{ID: id, Name: "bmw", Year: 2000, Brand: "bmw", Fuel: "petrol",
		Engine: model.Engine{EngineID: id1, Displacement: 20, NoOfCylinders: 5, Ranges: 0}}

	testcases := []struct {
		desc     string
		input    model.Car
		exOutput model.Car
		err      error
	}{
		{"error in store layer", car, model.Car{}, errors.New("error in store")},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		mockCar.EXPECT().CarUpdate(ctx, &tc.input).Return(&tc.exOutput, tc.err)
		resp, err := svc.Update(ctx, &tc.input)
		assert.Equal(t, resp, &tc.exOutput)
		assert.Equal(t, err, tc.err)
	}
}

// TestCarUpdateInvalidID is used to test the update method
func TestCarUpdateInvalidID(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	id1 := uuid.New()
	car := model.Car{ID: id, Name: "bmw", Year: 2000, Brand: "bmw", Fuel: "petrol",
		Engine: model.Engine{EngineID: id1, Displacement: 20, NoOfCylinders: 5, Ranges: 0}}

	testcases := []struct {
		desc     string
		input    model.Car
		exOutput model.Car
		carErr   error
		err      error
	}{
		{"error in id", car, model.Car{}, nil, errors.New("error in id")},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		mockCar.EXPECT().CarUpdate(ctx, &tc.input).Return(&tc.exOutput, tc.carErr)
		mockCar.EXPECT().CarGet(ctx, tc.input.ID).Return(&tc.exOutput, tc.err)

		resp, err := svc.Update(ctx, &tc.input)
		assert.Equal(t, resp, &tc.exOutput)
		assert.Equal(t, err, tc.err)
	}
}

//  TestUpdateEngine is used to test the update method which has  error in engine store layer
func TestUpdateEngineErr(t *testing.T) {
	ctx := context.Background()
	id1 := uuid.New()
	car := model.Car{ID: id1, Name: "bmw", Year: 2000, Brand: "bmw", Fuel: "petrol",
		Engine: model.Engine{EngineID: id1, Displacement: 20, NoOfCylinders: 5, Ranges: 0}}

	testcases := []struct {
		desc  string
		input model.Car

		exOutput model.Car
		carErr   error
		err      error
	}{
		{"error in engine store", car, model.Car{}, nil, errors.New("error in engine")},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		mockCar.EXPECT().CarUpdate(ctx, &tc.input).Return(&tc.exOutput, tc.carErr)
		mockCar.EXPECT().CarGet(ctx, tc.input.ID).Return(&tc.exOutput, tc.carErr)
		mockEngine.EXPECT().EngineUpdate(ctx, &tc.input.Engine).Return(&tc.exOutput.Engine, tc.err)

		resp, err := svc.Update(ctx, &tc.input)
		assert.Equal(t, resp, &tc.exOutput)
		assert.Equal(t, err, tc.err)
	}
}

// TestDelete is used to test the delete method
func TestDelete(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	testcases := []struct {
		desc   string
		id     uuid.UUID
		idErr  error
		CarErr error
		EngErr error
		err    error
	}{

		{"error from datastore layer", id, errors.New("error from datastore layer"),
			nil, nil, errors.New("error from datastore layer")},
	}
	car := model.Car{ID: id, Name: "bmw", Year: 2000, Brand: "bmw", Fuel: "electric",
		Engine: model.Engine{Displacement: 0, NoOfCylinders: 0, Ranges: 10}}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for i, tc := range testcases {
		mockCar.EXPECT().CarGet(context.TODO(), tc.id).Return(&car, tc.idErr)
		err := svc.Delete(ctx, tc.id)

		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("[Test %v] Failed. Got %v Expected %v", i+1, err, tc.err)
		}
	}
}

// TestDeleteSuccess is used to test the delete method
func TestDeleteSuccess(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	testcases := []struct {
		desc   string
		id     uuid.UUID
		idErr  error
		CarErr error
		EngErr error
		err    error
	}{
		{"success case", id, nil, nil, nil, nil},
		{"invalid id", id, nil, errors.New("invalid id"), nil, errors.New("invalid id")},
	}
	car := model.Car{ID: id, Name: "bmw", Year: 2000, Brand: "bmw", Fuel: "electric",
		Engine: model.Engine{Displacement: 0, NoOfCylinders: 0, Ranges: 10}}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for i, tc := range testcases {
		mockCar.EXPECT().CarGet(ctx, tc.id).Return(&car, tc.idErr)
		mockEngine.EXPECT().EngineDelete(ctx, car.Engine.EngineID).Return(tc.EngErr)
		mockCar.EXPECT().CarDelete(ctx, tc.id).Return(tc.CarErr)

		err := svc.Delete(ctx, tc.id)

		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("[Test %v] Failed. Got %v Expected %v", i+1, err, tc.err)
		}
	}
}

// TestDeleteEngError is used to test the delete method
func TestDeleteEngError(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	testcases := []struct {
		desc   string
		id     uuid.UUID
		idErr  error
		CarErr error
		EngErr error
		err    error
	}{
		{"invalid id", id, nil, nil, errors.New("invalid id"), errors.New("invalid id")},
	}
	car := model.Car{ID: id, Name: "bmw", Year: 2000, Brand: "bmw", Fuel: "electric",
		Engine: model.Engine{Displacement: 0, NoOfCylinders: 0, Ranges: 10}}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for i, tc := range testcases {
		mockCar.EXPECT().CarGet(context.TODO(), tc.id).Return(&car, tc.idErr)
		mockEngine.EXPECT().EngineDelete(ctx, car.Engine.EngineID).Return(tc.EngErr)

		err := svc.Delete(ctx, tc.id)

		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("[Test %v] Failed. Got %v Expected %v", i+1, err, tc.err)
		}
	}
}
