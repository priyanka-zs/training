package store

import (
	"context"
	"database/sql"
	"employee/models"
)

type DB struct {
	db *sql.DB
}

func New(d *sql.DB) DB {
	return DB{db: d}
}

func (d DB) CreateStore(ctx context.Context, emp models.Employee) (models.Employee, error) {
	_, err := d.db.ExecContext(ctx, "insert into employee(id,name,age)VALUES (?,?,?)", emp.Id, emp.Name, emp.Age)
	if err != nil {
		return models.Employee{}, err
	}
	return emp, nil
}
func (d DB) GetById(ctx context.Context, id int) (models.Employee, error) {
	row := d.db.QueryRowContext(ctx, "select * from employee where Id=?", id)
	var emp models.Employee
	err := row.Scan(&emp.Id, &emp.Name, &emp.Age)
	if err != nil {
		return models.Employee{}, err
	}
	return emp, nil
}
