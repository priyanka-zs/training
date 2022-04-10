package engine

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

func New(db *sql.DB) DB {
	return DB{db: db}
}

const (
	getByID = "SELECT Engine_Id,Displacement,No_of_cylinders,`Range` FROM engine where Engine_Id = ?"
	post    = "INSERT INTO engine(Engine_Id,Displacement,No_of_cylinders,`Range`) VALUES(?,?,?,?)"
	put     = "UPDATE engine SET Displacement=?,No_of_cylinders=?,`Range`=? where Engine_Id=?"
	del     = "delete from engine where Engine_Id=?"
)

// EngineCreate is used to insert a row into the engine table
func (d DB) EngineCreate(ctx context.Context, engine *model.Engine) (*model.Engine, error) {
	_, err := d.db.ExecContext(ctx, post, engine.EngineID, engine.Displacement, engine.NoOfCylinders, engine.Ranges)
	if err != nil {
		return &model.Engine{}, errors.New("error in insertion")
	}

	return engine, nil
}

// EngineGet method takes id as input and returns a row with the given id
func (d DB) EngineGet(ctx context.Context, id uuid.UUID) (*model.Engine, error) {
	var engine = model.Engine{}

	rows := d.db.QueryRowContext(ctx, getByID, id.String())

	err := rows.Scan(&engine.EngineID, &engine.Displacement, &engine.NoOfCylinders, &engine.Ranges)
	if err != nil {
		return &model.Engine{}, errors.New("query error")
	}

	return &engine, nil
}

// EngineUpdate method is used to update/modify a particular row
func (d DB) EngineUpdate(ctx context.Context, engine *model.Engine) (*model.Engine, error) {
	_, err := d.db.ExecContext(ctx, put, engine.Displacement, engine.NoOfCylinders, engine.Ranges, engine.EngineID)
	if err != nil {
		return &model.Engine{}, errors.New("query error")
	}

	return engine, nil
}

// EngineDelete is used to delete a row based on given id.It takes id as input.
func (d DB) EngineDelete(ctx context.Context, id uuid.UUID) error {
	_, err := d.db.ExecContext(ctx, del, id)
	if err != nil {
		return err
	}

	return nil
}
