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
	DeleteFarm(ctx context.Context, ID int) error
	IsFarmExists(ctx context.Context, farmName string, ID int) (bool, error)
	GetFarms(ctx context.Context) ([]farm.Farms, error)
	GetFarm(ctx context.Context, ID int) ([]farm.Farm, error)
	UpdateFarm(ctx context.Context, ID int, farmName string) error
}

func (fc FarmConfig) UpdateFarm(ctx context.Context, ID int, farmName string) error {

	script := `UPDATE farm set farm_name = $1 where id = $2`

	_, err := fc.db.Backend.Write.Exec(script, farmName, ID)
	if err != nil {
		return err
	}

	return nil
}

func (fc FarmConfig) GetFarm(ctx context.Context, ID int) ([]farm.Farm, error) {

	var res []farm.Farm

	script := `select 
	f.id, 
    p.id as pond_id,
	f.farm_name,
	p.pond_name,
    p.created_at
	from farm f 
	inner join pond p on p.farm_id = f.id
	where p.is_deleted = false and f.is_deleted = false and f.id = $1`

	query, args, err := fc.db.Backend.Read.In(script, ID)
	if err != nil {
		return nil, err
	}

	query = fc.db.Backend.Read.Rebind(query)
	err = fc.db.Backend.Read.Select(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return res, nil
}

func (fc FarmConfig) DeleteFarm(ctx context.Context, ID int) error {
	script := `UPDATE farm set is_deleted = true where id = $1`

	_, err := fc.db.Backend.Write.Exec(script, ID)
	if err != nil {
		return err
	}

	return nil
}

func (fc FarmConfig) GetFarms(ctx context.Context) ([]farm.Farms, error) {

	var res []farm.Farms

	script := `select f.id, f.farm_name, f.created_at from farm f where f.is_deleted = false`

	query, args, err := fc.db.Backend.Read.In(script)
	if err != nil {
		return res, err
	}

	query = fc.db.Backend.Read.Rebind(query)
	err = fc.db.Backend.Read.Select(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return res, err
	}
	return res, nil
}

func (fc FarmConfig) IsFarmExists(ctx context.Context, farmName string, ID int) (bool, error) {
	var isExist bool

	if farmName != "" {
		script := `select exists(select * from farm where farm_name = $1 and is_deleted = false)`

		query, args, err := fc.db.Backend.Read.In(script, farmName)
		if err != nil {
			return isExist, err
		}

		query = fc.db.Backend.Read.Rebind(query)
		err = fc.db.Backend.Read.GetContext(ctx, &isExist, query, args...)
		if err != nil && err != sql.ErrNoRows {
			return isExist, err
		}
	} else if ID != 0 {
		script := `select exists(select * from farm where id = $1 and is_deleted = false)`

		query, args, err := fc.db.Backend.Read.In(script, ID)
		if err != nil {
			return isExist, err
		}

		query = fc.db.Backend.Read.Rebind(query)
		err = fc.db.Backend.Read.GetContext(ctx, &isExist, query, args...)
		if err != nil && err != sql.ErrNoRows {
			return isExist, err
		}
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
