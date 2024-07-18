package farmService

import (
	"context"
	"errors"
	"time"

	"github.com/lukaslinardi/delos_aqua_api/domain/model/farm"
	"github.com/lukaslinardi/delos_aqua_api/domain/model/general"
	"github.com/lukaslinardi/delos_aqua_api/domain/model/pond"
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
	GetFarms(ctx context.Context) ([]farm.Farms, map[string]string, error)
	DeleteFarm(ctx context.Context, ID int) (map[string]string, error)
	GetFarm(ctx context.Context, ID int) (*farm.FarmRes, map[string]string, error)
}

func (fs FarmService) GetFarm(ctx context.Context, ID int) (*farm.FarmRes, map[string]string, error) {

	var res farm.FarmRes

	internalServerError := func(err error) (*farm.FarmRes, map[string]string, error) {
		return nil, map[string]string{
			"en": "Failed ! There's some trouble on our system, please try again",
			"id": "Gagal ! Terjadi kesalahan pada sistem, silahkan coba lagi",
		}, err
	}

    isExists, err := fs.db.Farm.IsFarmExists(ctx, "", ID)
    if err != nil {
		fs.log.WithField("request", utils.StructToString(isExists)).WithError(err).Errorf("failed to check farm")
		return nil, map[string]string{
			"en": "Failed ! There's some trouble on our system, please try again",
			"id": "Gagal ! Terjadi kesalahan pada sistem, silahkan coba lagi",
		}, errors.New("farm name already exists")
    }


    if !isExists {
		fs.log.WithField("request", utils.StructToString(isExists)).WithError(err).Errorf("farm not exists")
		return nil, map[string]string{
			"en": "farm not exists",
			"id": "farm tidak ada",
		}, errors.New("farm not exists")
    }

	data, err := fs.db.Farm.GetFarm(ctx, ID)
	if err != nil {
		fs.log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("failed to get farm detail")
		return internalServerError(err)
	}

	for _, value := range data {
		res.Ponds = append(res.Ponds, pond.Ponds{
			ID:        value.PondID,
			FarmID:    value.ID,
			PondName:  value.PondName,
			CreatedAt: value.CreatedAt,
		})
		res.ID = value.ID
		res.FarmName = value.FarmName
	}

	return &res, map[string]string{
		"en": "success",
		"id": "sukses",
	}, nil
}

func (fs FarmService) DeleteFarm(ctx context.Context, ID int) (map[string]string, error) {

	internalServerError := func(err error) (map[string]string, error) {
		return map[string]string{
			"en": "Failed ! There's some trouble on our system, please try again",
			"id": "Gagal ! Terjadi kesalahan pada sistem, silahkan coba lagi",
		}, err
	}

	isExists, err := fs.db.Farm.IsFarmExists(ctx, "", ID)
	if err != nil {
		fs.log.WithField("request", utils.StructToString(isExists)).WithError(err).Errorf("failed to check farm")
		return map[string]string{
			"en": "failed to check farm",
			"id": "gagal untuk cek farm",
		}, err
	}

	if !isExists {
		fs.log.WithField("request", utils.StructToString(isExists)).WithError(err).Errorf("farm name already exists")
		return map[string]string{
			"en": "farm not exists",
			"id": "farm tidak exists",
		}, errors.New("farm name already exists")
	}

	err = fs.db.Farm.DeleteFarm(ctx, ID)
	if err != nil {
		fs.log.WithField("request", utils.StructToString(err)).WithError(err).Errorf("failed to delete farm")
		return internalServerError(err)
	}
	return map[string]string{
		"en": "success",
		"id": "sukses",
	}, nil

}

func (fs FarmService) GetFarms(ctx context.Context) ([]farm.Farms, map[string]string, error) {

	internalServerError := func(err error) ([]farm.Farms, map[string]string, error) {
		return nil, map[string]string{
			"en": "Failed ! There's some trouble on our system, please try again",
			"id": "Gagal ! Terjadi kesalahan pada sistem, silahkan coba lagi",
		}, err
	}

	data, err := fs.db.Farm.GetFarms(ctx)
	if err != nil {
		fs.log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("Sign Up | Low | fail to begin transaction")
		return internalServerError(err)
	}

	return data, map[string]string{
		"en": "success",
		"id": "sukses",
	}, nil
}

func (fs FarmService) InsertFarm(ctx context.Context, data farm.InsertFarm) (map[string]string, error) {

	tx, err := fs.dbConn.Backend.Read.Begin()
	if err != nil {
		fs.log.WithField("request", utils.StructToString(tx)).WithError(err).Errorf("failed to start Tx")
		tx.Rollback()
		return map[string]string{
			"en": "failed to begin Tx",
			"id": "gagal memulai Tx",
		}, err
	}

	isExists, err := fs.db.Farm.IsFarmExists(ctx, data.FarmName, 0)
	if err != nil {
		fs.log.WithField("request", utils.StructToString(tx)).WithError(err).Errorf("failed to check farm")
		tx.Rollback()
		return map[string]string{
			"en": "failed to check farm",
			"id": "gagal untuk cek farm",
		}, err
	}

	if isExists {
		fs.log.WithField("request", utils.StructToString(tx)).WithError(err).Errorf("farm name already exists")
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
		fs.log.WithField("request", utils.StructToString(tx)).WithError(err).Errorf("failed to insert farm")
		tx.Rollback()
		return map[string]string{
			"en": "failed to insert farm",
			"id": "gagal memasukkan farm",
		}, err
	}

	err = tx.Commit()
	if err != nil {
		fs.log.WithField("request", utils.StructToString(tx)).WithError(err).Errorf("failed commit")
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
