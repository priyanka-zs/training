package car

import (
	"database/sql"
	_ "database/sql"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"gofr_cardealership/db"

	"gofr_cardealership/model"

	"github.com/google/uuid"
)

type car struct {
	db []model.Car
}

func New() *car {
	return &car{db.Car}
}

// CarCreate method is used to insert a row into the car table
func (c *car) CarCreate(ctx *gofr.Context, car *model.Car) (*model.Car, error) {
	car.ID = uuid.New()
	car.Engine.EngineID = uuid.New()
	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO Car(Id,Name,Year,Brand,FuelType,Engine_Id)"+
		"VALUES(?,?,?,?,?,?)", car.ID, car.Name, car.Year, car.Brand, car.Fuel, car.Engine.EngineID)

	if err != nil {
		return &model.Car{}, errors.Error("internal server error")
	}
	return car, nil
}

// CarGet is used to get a row of a car based on given id
func (c car) CarGet(ctx *gofr.Context, id uuid.UUID) (*model.Car, error) {
	var car = &model.Car{}

	rows := ctx.DB().QueryRowContext(ctx, "SELECT Id,Name,Year,Brand,FuelType,Engine_Id FROM Car where Id=?", id)
	err := rows.Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.Fuel, &car.Engine.EngineID)

	if err != nil {
		return &model.Car{}, errors.Error("internal server error")
	}

	return car, nil
}

// CarUpdate method is used to update/modify a particular row in car table
func (c car) CarUpdate(ctx *gofr.Context, car *model.Car) (*model.Car, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE Car SET Name=?,Year=?,Brand=?,FuelType=? "+
		"where Id=?", car.Name, car.Year, car.Brand, car.Fuel, car.ID)

	if err != nil {
		return &model.Car{}, errors.Error("internal server error")
	}

	return car, nil
}

// CarDelete method is used to delete a row in car table based on given id
func (c car) CarDelete(ctx *gofr.Context, id uuid.UUID) error {
	_, err := ctx.DB().ExecContext(ctx, "delete from Car where Id=?", id)

	if err != nil {
		return errors.Error("internal server error")
	}

	return nil
}

// CarGetByBrand method takes brand as input and returns rows with the given brand
func (c car) CarGetByBrand(ctx *gofr.Context, brand string, isEngine bool) ([]*model.Car, error) {
	rows, err := ctx.DB().QueryContext(ctx, "select * from Car where Brand=?", brand)
	if err != nil {
		return nil, errors.Error("internal server error")
	}

	if rows.Err() != nil {
		return []*model.Car{}, errors.Error("row error")
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var res []*model.Car

	for rows.Next() {
		var car model.Car

		err := rows.Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.Fuel, &car.Engine.EngineID)
		if err != nil {
			return []*model.Car{}, errors.Error("error while scanning")
		}

		res = append(res, &car)
	}

	return res, nil
}
