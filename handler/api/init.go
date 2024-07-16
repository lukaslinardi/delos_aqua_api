package api

import (
	"github.com/lukaslinardi/delos_aqua_api/domain/model/general"
	authHandler "github.com/lukaslinardi/delos_aqua_api/handler/api/auth"
	farmHandler "github.com/lukaslinardi/delos_aqua_api/handler/api/farm"
	pondHandler "github.com/lukaslinardi/delos_aqua_api/handler/api/pond"
	"github.com/lukaslinardi/delos_aqua_api/service"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Farm   farmHandler.FarmDataHandler
	Pond   pondHandler.PondDataHandler
	Public authHandler.PublicHandler
}

func NewHandler(sv service.Service, conf general.AppService, logger *logrus.Logger) Handler {
	return Handler{
		Farm:   farmHandler.NewFarmDataHandler(sv, conf, logger),
		Public: authHandler.NewPublicHandler(conf, logger),
		Pond:   pondHandler.NewPondDataHandler(sv, conf, logger),
	}

}
