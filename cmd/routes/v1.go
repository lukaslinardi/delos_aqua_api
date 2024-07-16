package routes

import (
	//"net/http"

	"net/http"

	"github.com/gorilla/mux"
	api "github.com/lukaslinardi/delos_aqua_api/handler"
)

func getV1(router, routerJWT *mux.Router, handler api.Handler) {
	router.HandleFunc("/v1/farm", handler.Farm.Farm.InsertFarm).Methods(http.MethodPost)
	router.HandleFunc("/v1/farms", handler.Farm.Farm.GetFarms).Methods(http.MethodGet)
	router.HandleFunc("/v1/delete-farm", handler.Farm.Farm.DeleteFarm).Methods(http.MethodDelete)

	router.HandleFunc("/v1/pond", handler.Pond.Pond.InsertPond).Methods(http.MethodPost)
	router.HandleFunc("/v1/ponds", handler.Pond.Pond.GetPonds).Methods(http.MethodGet)
	router.HandleFunc("/v1/delete-pond", handler.Pond.Pond.DeletePond).Methods(http.MethodDelete)
	// router.HandleFunc("/v1/signup", handler.Auth.Auth.SignUp).Methods(http.MethodPost)
	// router.HandleFunc("/v1/login", handler.Auth.Auth.Login).Methods(http.MethodPost)
	// router.HandleFunc("/v1/psef/forget-password", handler.Auth.Auth.ForgetPassword).Methods(http.MethodGet)
	// routerJWT.HandleFunc("/v1/psef/outlet", handler.User.User.GetOutletList).Methods(http.MethodGet)
	// routerJWT.HandleFunc("/v1/psef/outlet-detail", handler.User.User.GetOutletDetail).Methods(http.MethodGet)
	// routerJWT.HandleFunc("/v1/psef/product", handler.User.User.GetProductList).Methods(http.MethodGet)
	// routerJWT.HandleFunc("/v1/psef/product-detail", handler.User.User.GetProductDetail).Methods(http.MethodGet)
}
