package db

import (
	"context"
	"database/sql"

	"github.com/lukaslinardi/delos_aqua_api/domain/model/farm"
	"github.com/lukaslinardi/delos_aqua_api/infra"
	"github.com/sirupsen/logrus"
)

type FarmConfig struct {
	db  *infra.DatabaseList
	log *logrus.Logger
}

func newFarm(db *infra.DatabaseList, logger *logrus.Logger) FarmConfig {
	return FarmConfig{
		db:  db,
		log: logger,
	}
}

type Farm interface {
	InsertFarm(ctx context.Context, tx *sql.Tx, data farm.InsertFarm) error
	IsFarmExists(ctx context.Context, farmName string) (bool, error)
}

func (fc FarmConfig) IsFarmExists(ctx context.Context, farmName string) (bool, error) {
	var isExist bool

	script := `select exists(select * from farm where farm_name = $1)`

	query, args, err := fc.db.Backend.Read.In(script, farmName)
	if err != nil {
		return isExist, err
	}

	query = fc.db.Backend.Read.Rebind(query)
	err = fc.db.Backend.Read.GetContext(ctx, &isExist, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return isExist, err
	}
	return isExist, nil
}

func (fc FarmConfig) InsertFarm(ctx context.Context, tx *sql.Tx, data farm.InsertFarm) error {
	script := `INSERT INTO farm (farm_name, is_deleted, created_at, updated_at)
	VALUES($1, $2, $3, $4);`

	param := make([]interface{}, 0)

	param = append(param, data.FarmName)
	param = append(param, data.IsDeleted)
	param = append(param, data.CreatedAt)
	param = append(param, data.UpdatedAt)

	query, args, err := fc.db.Backend.Read.In(script, param...)

	query = fc.db.Backend.Read.Rebind(query)

	var res *sql.Row
	if tx == nil {
		res = fc.db.Backend.Write.QueryRow(ctx, query, args...)
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
