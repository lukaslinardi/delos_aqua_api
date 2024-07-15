package api

import (
	"github.com/lukaslinardi/delos_aqua_api/domain/model/general"
	farmHandler "github.com/lukaslinardi/delos_aqua_api/handler/api/farm"
	authHandler "github.com/lukaslinardi/delos_aqua_api/handler/api/auth"
	"github.com/lukaslinardi/delos_aqua_api/service"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Farm farmHandler.FarmDataHandler
    Public authHandler.PublicHandler
}

func NewHandler(sv service.Service, conf general.AppService, logger *logrus.Logger) Handler {
	return Handler{
		Farm: farmHandler.NewFarmDataHandler(sv, conf, logger),
		Public: authHandler.NewPublicHandler(conf, logger),
	}

}
