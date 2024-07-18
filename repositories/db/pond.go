package db

import (
	"context"
	"database/sql"

	"github.com/lukaslinardi/delos_aqua_api/domain/model/pond"
	"github.com/lukaslinardi/delos_aqua_api/infra"
	"github.com/sirupsen/logrus"
)

type PondConfig struct {
	db  *infra.DatabaseList
	log *logrus.Logger
}

func newPond(db *infra.DatabaseList, logger *logrus.Logger) PondConfig {
	return PondConfig{
		db:  db,
		log: logger,
	}
}

type Pond interface {
	InsertPond(ctx context.Context, tx *sql.Tx, data pond.InsertPond) error
	DeletePond(ctx context.Context, ID int) error
	IsPondExists(ctx context.Context, pondName string, ID int) (bool, error)
	GetPonds(ctx context.Context) ([]pond.Ponds, error)
	GetPond(ctx context.Context, ID int) (*pond.Pond, error)
}

func (pc PondConfig) GetPond(ctx context.Context, ID int) (*pond.Pond, error) {
	var res pond.Pond

	script := `select p.id, 
	   f.id as farm_id,
	   f.farm_name,
	   p.pond_name,
       p.created_at
	   from pond p
	   inner join farm f on f.id = p.farm_id where p.id = $1 and p.is_deleted = false and f.is_deleted = false`

	query, args, err := pc.db.Backend.Read.In(script, ID)
	if err != nil {
		return &res, err
	}

	query = pc.db.Backend.Read.Rebind(query)
	err = pc.db.Backend.Read.GetContext(ctx, &res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return &res, err
	}
	return &res, nil

}

func (pc PondConfig) DeletePond(ctx context.Context, ID int) error {
	script := `UPDATE pond set is_deleted = true where id = $1`

	_, err := pc.db.Backend.Write.Exec(script, ID)
	if err != nil {
		return err
	}

	return nil
}

func (pc PondConfig) GetPonds(ctx context.Context) ([]pond.Ponds, error) {
	var res []pond.Ponds

	script := `select p.id, p.pond_name, p.farm_id, p.created_at from pond p where is_deleted = false`

	query, args, err := pc.db.Backend.Read.In(script)
	if err != nil {
		return res, err
	}

	query = pc.db.Backend.Read.Rebind(query)
	err = pc.db.Backend.Read.Select(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return res, err
	}
	return res, nil
}

func (pc PondConfig) IsPondExists(ctx context.Context, pondName string, ID int) (bool, error) {

	var isExist bool

	if pondName != "" {
		script := `select exists(select * from pond where pond_name = $1 and is_deleted = false)`

		query, args, err := pc.db.Backend.Read.In(script, pondName)
		if err != nil {
			return isExist, err
		}

		query = pc.db.Backend.Read.Rebind(query)
		err = pc.db.Backend.Read.GetContext(ctx, &isExist, query, args...)
		if err != nil && err != sql.ErrNoRows {
			return isExist, err
		}
	} else if ID != 0 {
		script := `select exists(select * from pond where id = $1 and is_deleted = false)`

		query, args, err := pc.db.Backend.Read.In(script, ID)
		if err != nil {
			return isExist, err
		}

		query = pc.db.Backend.Read.Rebind(query)
		err = pc.db.Backend.Read.GetContext(ctx, &isExist, query, args...)
		if err != nil && err != sql.ErrNoRows {
			return isExist, err
		}
	}

	return isExist, nil
}

func (pc PondConfig) InsertPond(ctx context.Context, tx *sql.Tx, data pond.InsertPond) error {
	script := `INSERT INTO pond (pond_name, farm_id, is_deleted, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5);`

	param := make([]interface{}, 0)

	param = append(param, data.PondName)
	param = append(param, data.FarmID)
	param = append(param, data.IsDeleted)
	param = append(param, data.CreatedAt)
	param = append(param, data.UpdatedAt)

	query, args, err := pc.db.Backend.Read.In(script, param...)

	query = pc.db.Backend.Read.Rebind(query)

	var res *sql.Row
	if tx == nil {
		res = pc.db.Backend.Write.QueryRow(ctx, query, args...)
	} else {
		res = tx.QueryRowContext(ctx, query, args...)
	}

	if err != nil {
		return err
	}

	err = res.Err()
	if err != nil {
		return err
	}

	return nil
}
