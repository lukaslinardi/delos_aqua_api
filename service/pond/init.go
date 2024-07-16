package pondService

import (
	"github.com/lukaslinardi/delos_aqua_api/domain/model/general"
	"github.com/lukaslinardi/delos_aqua_api/infra"
	repository "github.com/lukaslinardi/delos_aqua_api/repositories"
	"github.com/sirupsen/logrus"
)

type PondData struct {
	Pond PondService
}

func NewPond(repo repository.Repo, conf general.AppService, dbList *infra.DatabaseList, logger *logrus.Logger) PondData {
	return PondData{
		Pond: newPondService(repo.Database, conf, dbList, logger),
	}
}
