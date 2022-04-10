package stores

import (
	"context"
	"database/sql"
	"errors"
	"petstore/models"

	"github.com/google/uuid"
)

type DB struct {
	db *sql.DB
}

func New(d *sql.DB) DB {
	return DB{db: d}
}

// Post Method is used to insert data into the database
func (d DB) Post(ctx context.Context, pet models.Pet) (models.Pet, error) {
	pet.ID = uuid.New()
	_, err := d.db.ExecContext(ctx, "insert into pet(Id,Name,Age,Species,Status)VALUES (?,?,?,?,?)",
		pet.ID, pet.Name, pet.Age, pet.Species, pet.Status)

	if err != nil {
		return models.Pet{}, errors.New("error in store layer")
	}

	return pet, nil
}

// GetByID takes id as input and returns a particular row as output
func (d DB) GetByID(ctx context.Context, id uuid.UUID) (models.Pet, error) {
	row := d.db.QueryRowContext(ctx, "select * from pet where id=?", id)

	var pet models.Pet

	err := row.Scan(&pet.ID, &pet.Name, &pet.Age, &pet.Species, &pet.Status)

	if err != nil {
		return models.Pet{}, errors.New("error in store layer")
	}

	return pet, nil
}

func (d DB) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := d.db.ExecContext(ctx, "delete from pet where Id=?", id)
	if err != nil {
		return err
	}

	return nil
}
