package engine

import (
	_ "database/sql"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/google/uuid"
	"gofr_cardealership/db"

	"gofr_cardealership/model"
)

type engine struct {
	db []model.Engine
}

func New() *engine {
	return &engine{db.Engine}
}

const (
	getByID = "SELECT EngineID,Displacement,NoOfCylinders,Ranges FROM Engine where EngineID = ?"
	post    = "INSERT INTO Engine(EngineID,Displacement,NoOfCylinders,Ranges) VALUES(?,?,?,?)"
	put     = "UPDATE Engine SET Displacement=?,NoOfCylinders=?,Ranges=? where EngineID=?"
	del     = "delete from Engine where EngineID=?"
)

// EngineCreate is used to insert a row into the engine table
func (e *engine) EngineCreate(ctx *gofr.Context, engine *model.Engine) (*model.Engine, error) {
	_, err := ctx.DB().ExecContext(ctx, post, engine.EngineID, engine.Displacement, engine.NoOfCylinders, engine.Ranges)
	if err != nil {
		return &model.Engine{}, errors.Error("Internal server error")
	}

	return engine, nil
}

// EngineGet method takes id as input and returns a row with the given id
func (e *engine) EngineGet(ctx *gofr.Context, id uuid.UUID) (*model.Engine, error) {
	var engine = model.Engine{}

	rows := ctx.DB().QueryRowContext(ctx, getByID, id.String())
	err := rows.Scan(&engine.EngineID, &engine.Displacement, &engine.NoOfCylinders, &engine.Ranges)
	if err != nil {
		return &model.Engine{}, errors.Error("Internal server error")
	}

	return &engine, nil
}

// EngineUpdate method is used to update/modify a particular row
func (e *engine) EngineUpdate(ctx *gofr.Context, engine *model.Engine) (*model.Engine, error) {
	_, err := ctx.DB().Exec(put, engine.Displacement, engine.NoOfCylinders, engine.Ranges, engine.EngineID)
	if err != nil {
		return &model.Engine{}, errors.Error("query error")
	}

	return engine, nil
}

// EngineDelete is used to delete a row based on given id.It takes id as input.
func (e *engine) EngineDelete(ctx *gofr.Context, id uuid.UUID) error {
	_, err := ctx.DB().ExecContext(ctx, del, id)
	if err != nil {
		return errors.Error("query error")
	}

	return nil
}
