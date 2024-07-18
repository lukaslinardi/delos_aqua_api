package pondService

import (
	"context"
	"errors"
	"time"

	"github.com/lukaslinardi/delos_aqua_api/domain/model/general"
	"github.com/lukaslinardi/delos_aqua_api/domain/model/pond"
	"github.com/lukaslinardi/delos_aqua_api/domain/utils"
	"github.com/lukaslinardi/delos_aqua_api/infra"
	"github.com/lukaslinardi/delos_aqua_api/repositories/db"
	"github.com/sirupsen/logrus"
)

type PondService struct {
	db     db.Database
	conf   general.AppService
	dbConn *infra.DatabaseList
	log    *logrus.Logger
}

func newPondService(db db.Database, conf general.AppService, dbConn *infra.DatabaseList, logger *logrus.Logger) PondService {
	return PondService{
		db:     db,
		conf:   conf,
		dbConn: dbConn,
		log:    logger,
	}
}

type Pond interface {
	InsertPond(ctx context.Context, data pond.InsertPond) (map[string]string, error)
	DeletePond(ctx context.Context, ID int) (map[string]string, error)
	GetPonds(ctx context.Context) ([]pond.Ponds, map[string]string, error)
	GetPond(ctx context.Context, ID int) (*pond.Pond, map[string]string, error)
}

func (ps PondService) GetPond(ctx context.Context, ID int) (*pond.Pond, map[string]string, error) {

	internalServerError := func(err error) (*pond.Pond, map[string]string, error) {
		return nil, map[string]string{
			"en": "Failed ! There's some trouble on our system, please try again",
			"id": "Gagal ! Terjadi kesalahan pada sistem, silahkan coba lagi",
		}, err
	}

	data, err := ps.db.Pond.GetPond(ctx, ID)
	if err != nil {
		ps.log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("failed to get farm detail")
		return internalServerError(err)
	}

	return data, map[string]string{
		"en": "success",
		"id": "sukses",
	}, nil
}

func (ps PondService) DeletePond(ctx context.Context, ID int) (map[string]string, error) {

	internalServerError := func(err error) (map[string]string, error) {
		return map[string]string{
			"en": "Failed ! There's some trouble on our system, please try again",
			"id": "Gagal ! Terjadi kesalahan pada sistem, silahkan coba lagi",
		}, err
	}

	isExists, err := ps.db.Pond.IsPondExists(ctx, "", ID)
	if err != nil {
		ps.log.WithField("request", utils.StructToString(isExists)).WithError(err).Errorf("failed to check farm")
		return map[string]string{
			"en": "failed to check pond",
			"id": "gagal untuk cek pond",
		}, err
	}

	if !isExists {
		ps.log.WithField("request", utils.StructToString(isExists)).WithError(err).Errorf("farm name already exists")
		return map[string]string{
			"en": "pond not exists",
			"id": "pond tidak exists",
		}, errors.New("pond not exists")
	}

	err = ps.db.Pond.DeletePond(ctx, ID)
	if err != nil {
		ps.log.WithField("request", utils.StructToString(err)).WithError(err).Errorf("failed to delete farm")
		return internalServerError(err)
	}
	return map[string]string{
		"en": "success",
		"id": "sukses",
	}, nil

}

func (ps PondService) GetPonds(ctx context.Context) ([]pond.Ponds, map[string]string, error) {

	internalServerError := func(err error) ([]pond.Ponds, map[string]string, error) {
		return nil, map[string]string{
			"en": "Failed ! There's some trouble on our system, please try again",
			"id": "Gagal ! Terjadi kesalahan pada sistem, silahkan coba lagi",
		}, err
	}

	data, err := ps.db.Pond.GetPonds(ctx)
	if err != nil {
		ps.log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("failed to get ponds")
		return internalServerError(err)
	}

	return data, map[string]string{
		"en": "success",
		"id": "sukses",
	}, nil

}

func (ps PondService) InsertPond(ctx context.Context, data pond.InsertPond) (map[string]string, error) {

	tx, err := ps.dbConn.Backend.Read.Begin()
	if err != nil {
		ps.log.WithField("request", utils.StructToString(tx)).WithError(err).Errorf("fail to Start Tx")
		tx.Rollback()
		return map[string]string{
			"en": "failed to begin Tx",
			"id": "gagal memulai Tx",
		}, err
	}

	isExists, err := ps.db.Farm.IsFarmExists(ctx, "", data.FarmID)
	if err != nil {
		ps.log.WithField("request", utils.StructToString(tx)).WithError(err).Errorf("fail to check Farm")
		tx.Rollback()
		return map[string]string{
			"en": "failed to check farm",
			"id": "gagal untuk cek farm",
		}, err
	}

	if !isExists {
		ps.log.WithField("request", utils.StructToString(tx)).WithError(err).Errorf("farm not exists")
		tx.Rollback()
		return map[string]string{
			"en": "farm not exists",
			"id": "farm tidak exists",
		}, errors.New("farm not exists")
	}

	isPondExists, err := ps.db.Pond.IsPondExists(ctx, data.PondName, 0)
	if isPondExists {
		ps.log.WithField("request", utils.StructToString(tx)).WithError(err).Errorf("pond name exits")
		tx.Rollback()
		return map[string]string{
			"en": "pond already exists",
			"id": "pond sudah ada",
		}, errors.New("pond already exists")
	}

	req := pond.InsertPond{
		IsDeleted: false,
		PondName:  data.PondName,
		FarmID:    data.FarmID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = ps.db.Pond.InsertPond(ctx, tx, req)
	if err != nil {
		ps.log.WithField("request", utils.StructToString(tx)).WithError(err).Errorf("failed to insert pond")
		tx.Rollback()
		return map[string]string{
			"en": "failed to insert farm",
			"id": "gagal memasukkan farm",
		}, err
	}

	err = tx.Commit()
	if err != nil {
		ps.log.WithField("request", utils.StructToString(tx)).WithError(err).Errorf("failed commit tx")
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
