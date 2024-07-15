package farmService

import (
	"context"
	"errors"
	"time"

	"github.com/lukaslinardi/delos_aqua_api/domain/model/farm"
	"github.com/lukaslinardi/delos_aqua_api/domain/model/general"
	"github.com/lukaslinardi/delos_aqua_api/domain/utils"
	"github.com/lukaslinardi/delos_aqua_api/infra"
	"github.com/lukaslinardi/delos_aqua_api/repositories/db"
	"github.com/sirupsen/logrus"
)

type FarmService struct {
	db     db.Database
	conf   general.AppService
	dbConn *infra.DatabaseList
	log    *logrus.Logger
}

func newFarmService(db db.Database, conf general.AppService, dbConn *infra.DatabaseList, logger *logrus.Logger) FarmService {
	return FarmService{
		db:     db,
		conf:   conf,
		dbConn: dbConn,
		log:    logger,
	}
}

type Farm interface {
	InsertFarm(ctx context.Context, data farm.InsertFarm) (map[string]string, error)
}

func (fs FarmService) InsertFarm(ctx context.Context, data farm.InsertFarm) (map[string]string, error) {

	tx, err := fs.dbConn.Backend.Read.Begin()
	if err != nil {
		fs.log.WithField("request", utils.StructToString(tx)).WithError(err).Errorf("Sign Up | Low | fail to begin transaction")
		tx.Rollback()
		return map[string]string{
			"en": "failed to begin Tx",
			"id": "gagal memulai Tx",
		}, err
	}

	isExists, err := fs.db.Farm.IsFarmExists(ctx, data.FarmName)
	if err != nil {
		fs.log.WithField("request", utils.StructToString(tx)).WithError(err).Errorf("Sign Up | Low | fail to begin transaction")
		tx.Rollback()
		return map[string]string{
			"en": "farm name already exists",
			"id": "nama farm sudah ada",
		}, errors.New("farm name already exists")
	}

	if isExists {
		fs.log.WithField("request", utils.StructToString(tx)).WithError(err).Errorf("Sign Up | Low | fail to begin transaction")
		tx.Rollback()
		return map[string]string{
			"en": "farm name already exists",
			"id": "nama farm sudah ada",
		}, errors.New("farm name already exists")
	}

	req := farm.InsertFarm{
		IsDeleted: false,
		FarmName:  data.FarmName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = fs.db.Farm.InsertFarm(ctx, tx, req)
	if err != nil {
		fs.log.WithField("request", utils.StructToString(tx)).WithError(err).Errorf("Sign Up | Low | fail to begin transaction")
		tx.Rollback()
		return map[string]string{
			"en": "failed to insert farm",
			"id": "gagal memasukkan farm",
		}, err
	}

	err = tx.Commit()
	if err != nil {
		fs.log.WithField("request", utils.StructToString(tx)).WithError(err).Errorf("Sign Up | Low | fail to begin transaction")
		tx.Rollback()
		return map[string]string{
			"en": "failed commit",
			"id": "gagal commit",
		}, err
	}

	return map[string]string{
		"en": "success",
		"id": "sukses",
	}, nil
}
