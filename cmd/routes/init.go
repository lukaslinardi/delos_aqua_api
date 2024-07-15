package routes

import (
	"github.com/gorilla/mux"
	//	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/domain/model/general"
	api "github.com/lukaslinardi/delos_aqua_api/handler"
	// "github.com/sirupsen/logrus"
)

func GetCoreEndpoint(handler api.Handler) *mux.Router {
	parentRoute := mux.NewRouter()

	jwtRoute := parentRoute.PathPrefix("").Subrouter()
	nonJWTRoute := parentRoute.PathPrefix("").Subrouter()

	// Middleware for public API
	nonJWTRoute.Use(handler.Public.AuthValidator)

	// Middleware
	// jwtRoute.Use(handler.Token.JWTValidator)

	// Get Endpoint.
	getV1(nonJWTRoute, jwtRoute, handler)

	return parentRoute
}
