package service

import (
	"github.com/lukaslinardi/delos_aqua_api/domain/model/general"
	"github.com/lukaslinardi/delos_aqua_api/infra"
	repository "github.com/lukaslinardi/delos_aqua_api/repositories"
	farmService "github.com/lukaslinardi/delos_aqua_api/service/farm"
	pondService "github.com/lukaslinardi/delos_aqua_api/service/pond"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Farm farmService.FarmData
	Pond pondService.PondData
}

func NewService(repo repository.Repo, conf general.AppService, dbList *infra.DatabaseList, logger *logrus.Logger) Service {
	return Service{
		Farm: farmService.NewFarm(repo, conf, dbList, logger),
		Pond: pondService.NewPond(repo, conf, dbList, logger),
	}
}
