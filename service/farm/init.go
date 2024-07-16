package farmService

import (
	"github.com/lukaslinardi/delos_aqua_api/domain/model/general"
	"github.com/lukaslinardi/delos_aqua_api/infra"
	repository "github.com/lukaslinardi/delos_aqua_api/repositories"
	"github.com/sirupsen/logrus"
)

type FarmData struct {
	Farm FarmService
}

func NewFarm(repo repository.Repo, conf general.AppService, dbList *infra.DatabaseList, logger *logrus.Logger) FarmData {
	return FarmData{
		Farm: newFarmService(repo.Database, conf, dbList, logger),
	}
}
