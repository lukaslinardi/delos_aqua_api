package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/lukaslinardi/delos_aqua_api/cmd/routes"
	mg "github.com/lukaslinardi/delos_aqua_api/domain/model/general"
	api "github.com/lukaslinardi/delos_aqua_api/handler"
	"github.com/lukaslinardi/delos_aqua_api/infra"
	repository "github.com/lukaslinardi/delos_aqua_api/repositories"
	"github.com/lukaslinardi/delos_aqua_api/service"
	"github.com/sirupsen/logrus"

)

func main() {

	conf, err := getConfigKey()
	if err != nil {
		panic(err)
	}

	handler, log, err := newRepoContext(conf)
	if err != nil {
		panic(err)
	}

	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	credentials := handlers.AllowCredentials()

	router := routes.GetCoreEndpoint(handler)

	// crsfProtect := csrf.Protect([]byte("test"))
	// csrf.Secure(false)

	port := fmt.Sprintf(":%s", "8080")
	log.Info("server listen to port ", port)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins, credentials)(router)))

}

func getConfigKey() (*mg.AppService, error) {

	conf := mg.AppService{
		Database: mg.Database{
			Read: mg.DBDetail{
				Username: "",
			},
			Write: mg.DBDetail{
				Username: "",
			},
		},
	}

	return &conf, nil

}

func newRepoContext(conf *mg.AppService) (api.Handler, *logrus.Logger, error) {
	var handler api.Handler

	// Init log
	logger := infra.NewLogger(conf)

	// Init DB Read Connection.
	dbRead := infra.NewDB(logger)
	dbRead.ConnectDB()
	if dbRead.Err != nil {
		return handler, logger, dbRead.Err
	}

	// Init DB Write Connection.
	dbWrite := infra.NewDB(logger)
	dbWrite.ConnectDB()
	if dbWrite.Err != nil {
		return handler, logger, dbWrite.Err
	}

	dbList := &infra.DatabaseList{
		Backend: infra.DatabaseType{
			Read:  &dbRead,
			Write: &dbWrite,
		},
	}

	repo := repository.NewRepo(dbList, *conf, logger)
	usecase := service.NewService(repo, *conf, dbList, logger)
	handler = api.NewHandler(usecase, *conf, logger)

	return handler, logger, nil
}
