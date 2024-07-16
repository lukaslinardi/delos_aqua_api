package farmHandler

import (
	// "encoding/json"
	// "io/ioutil"
	// "net/http"

	"encoding/json"
	"io/ioutil"
	"net/http"

	cg "github.com/lukaslinardi/delos_aqua_api/domain/constants/general"
	fm "github.com/lukaslinardi/delos_aqua_api/domain/model/farm"

	"github.com/lukaslinardi/delos_aqua_api/domain/model/general"
	"github.com/lukaslinardi/delos_aqua_api/domain/utils"
	"github.com/sirupsen/logrus"

	farmService "github.com/lukaslinardi/delos_aqua_api/service/farm"
)

type FarmHandler struct {
	Farm farmService.Farm
	conf general.AppService
	log  *logrus.Logger
}

func NewFarmHandler(farm farmService.Farm, conf general.AppService, logger *logrus.Logger) FarmHandler {
	return FarmHandler{
		Farm: farm,
		conf: conf,
		log:  logger,
	}
}

func (fh FarmHandler) GetFarms(res http.ResponseWriter, req *http.Request) {
	respData := &utils.ResponseDataV3{
		Status: cg.Fail,
	}

	data, message, err := fh.Farm.GetFarms(req.Context())
	if err != nil {
		respData.Message = message
		respData.ErrorDebug = err.Error()
		respData.ResponseFormatter()
		utils.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &utils.ResponseDataV3{
		Status:  cg.Success,
		Message: message,
		Detail:  data,
	}

	utils.WriteResponse(res, respData, http.StatusOK)
	return

}

func (fh FarmHandler) InsertFarm(res http.ResponseWriter, req *http.Request) {
	respData := &utils.ResponseDataV3{
		Status: cg.Fail,
	}

	var request fm.InsertFarm

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respData.Message = map[string]string{
			"en": cg.HandlerErrorRequestDataEmpty,
			"id": cg.HandlerErrorRequestDataEmptyID,
		}
		respData.ErrorDebug = err.Error()
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		respData.Message = map[string]string{
			"en": cg.HandlerErrorRequestDataNotValid,
			"id": cg.HandlerErrorRequestDataNotValidID,
		}
		respData.ErrorDebug = err.Error()
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	messages, err := fh.Farm.InsertFarm(req.Context(), request)
	if err != nil {
		respData.Message = messages
		respData.ErrorDebug = err.Error()
		respData.ResponseFormatter()
		utils.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &utils.ResponseDataV3{
		Status:  cg.Success,
		Message: messages,
	}

	utils.WriteResponse(res, respData, http.StatusOK)
	return

}
