package pondHandler

import (
	"github.com/lukaslinardi/delos_aqua_api/domain/model/general"
	"github.com/lukaslinardi/delos_aqua_api/service"
	"github.com/sirupsen/logrus"
)

type PondDataHandler struct {
	Pond PondHandler
}

func NewPondDataHandler(sv service.Service, conf general.AppService, logger *logrus.Logger) PondDataHandler {
	return PondDataHandler{
		Pond: NewPondHandler(sv.Pond.Pond, conf, logger),
	}
}
