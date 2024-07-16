package pondService

import (
	"github.com/lukaslinardi/delos_aqua_api/domain/model/general"
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

type Pond interface {}
