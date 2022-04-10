package car

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/zopsmart/GoLang-Interns-2022/model"
)

type DB struct {
	db *sql.DB
}

func New(d *sql.DB) DB {
	return DB{db: d}
}

// CarCreate method is used to insert a row into the car table
func (d DB) CarCreate(ctx context.Context, car *model.Car) (*model.Car, error) {
	car.ID = uuid.New()
	car.Engine.EngineID = uuid.New()
	_, err := d.db.ExecContext(ctx, "INSERT INTO car(Id,Name,Year,Brand,FuelType,Engine_Id)"+
		"VALUES(?,?,?,?,?,?)", car.ID, car.Name, car.Year, car.Brand, car.Fuel, car.Engine.EngineID)

	if err != nil {
		return &model.Car{}, errors.New("error in insertion")
	}
	return car, nil
}

// CarGet is used to get a row of a car based on given id
func (d DB) CarGet(ctx context.Context, id uuid.UUID) (*model.Car, error) {
	var car = &model.Car{}

	rows := d.db.QueryRowContext(ctx, "SELECT Id,Name,Year,Brand,FuelType,Engine_Id FROM car where Id=?", id)
	err := rows.Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.Fuel, &car.Engine.EngineID)

	if err != nil {
		return &model.Car{}, errors.New("error in get")
	}

	return car, nil
}

// CarUpdate method is used to update/modify a particular row in car table
func (d DB) CarUpdate(ctx context.Context, car *model.Car) (*model.Car, error) {
	_, err := d.db.ExecContext(ctx, "UPDATE car SET Name=?,Year=?,Brand=?,FuelType=? "+
		"where Id=?", car.Name, car.Year, car.Brand, car.Fuel, car.ID)

	if err != nil {
		return &model.Car{}, errors.New("id not in db")
	}

	return car, nil
}

// CarDelete method is used to delete a row in car table based on given id
func (d DB) CarDelete(ctx context.Context, id uuid.UUID) error {
	_, err := d.db.ExecContext(ctx, "delete from car where Id=?", id)

	if err != nil {
		return errors.New("error in deletion")
	}

	return nil
}

// CarGetByBrand method takes brand as input and returns rows with the given brand
func (d DB) CarGetByBrand(ctx context.Context, brand string, isEngine bool) ([]*model.Car, error) {
	rows, err := d.db.QueryContext(ctx, "select * from car where Brand=?", brand)
	if err != nil {
		return nil, errors.New("query error")
	}

	if rows.Err() != nil {
		return []*model.Car{}, errors.New("no rows fetched")
	}

	defer rows.Close()

	var res []*model.Car

	for rows.Next() {
		var car model.Car

		err := rows.Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.Fuel, &car.Engine.EngineID)
		if err != nil {
			return []*model.Car{}, errors.New("error while scanning")
		}

		res = append(res, &car)
	}

	return res, nil
}
