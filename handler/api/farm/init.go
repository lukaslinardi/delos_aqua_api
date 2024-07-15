package farmHandler

import (
	"github.com/lukaslinardi/delos_aqua_api/domain/model/general"
	"github.com/lukaslinardi/delos_aqua_api/service"
	"github.com/sirupsen/logrus"
)

type FarmDataHandler struct {
	Farm FarmHandler
}

func NewFarmDataHandler(sv service.Service, conf general.AppService, logger *logrus.Logger) FarmDataHandler {
	return FarmDataHandler{
		Farm: NewFarmHandler(sv.Farm.Farm, conf, logger),
	}
}
