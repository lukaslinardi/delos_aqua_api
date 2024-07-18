package farmHandler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

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

func (fh FarmHandler) UpdateFarm(res http.ResponseWriter, req *http.Request) {

	respData := &utils.ResponseDataV3{
		Status: cg.Fail,
	}

	ID := req.URL.Query().Get("ID")
	farmName := req.URL.Query().Get("farm_name")

	id, err := strconv.Atoi(ID)
	if err != nil {
		respData := &utils.ResponseDataV3{
			Status: cg.Fail,
			Message: map[string]string{
				"en": cg.HandlerErrorRequestDataNotValid,
				"id": cg.HandlerErrorRequestDataNotValidID,
			},
		}
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	message, err := fh.Farm.UpdateFarm(req.Context(), id, farmName)
	if err != nil {
		respData := &utils.ResponseDataV3{
			Status:  cg.Fail,
			Message: message,
		}
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	respData = &utils.ResponseDataV3{
		Status:  cg.Success,
		Message: message,
	}

	utils.WriteResponse(res, respData, http.StatusOK)
	return
}

func (fh FarmHandler) GetFarm(res http.ResponseWriter, req *http.Request) {
	respData := &utils.ResponseDataV3{
		Status: cg.Fail,
	}

	ID, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		respData := &utils.ResponseDataV3{
			Status: cg.Fail,
			Message: map[string]string{
				"en": cg.HandlerErrorRequestDataNotValid,
				"id": cg.HandlerErrorRequestDataNotValidID,
			},
		}
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	data, message, err := fh.Farm.GetFarm(req.Context(), ID)
	if err != nil {
		respData.Message = message
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

func (fh FarmHandler) DeleteFarm(res http.ResponseWriter, req *http.Request) {

	respData := &utils.ResponseDataV3{
		Status: cg.Fail,
	}

	ID := req.URL.Query().Get("ID")

	id, err := strconv.Atoi(ID)
	if err != nil {
		respData := &utils.ResponseDataV3{
			Status: cg.Fail,
			Message: map[string]string{
				"en": cg.HandlerErrorRequestDataNotValid,
				"id": cg.HandlerErrorRequestDataNotValidID,
			},
		}
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	message, err := fh.Farm.DeleteFarm(req.Context(), id)
	if err != nil {
		respData := &utils.ResponseDataV3{
			Status:  cg.Fail,
			Message: message,
		}
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	respData = &utils.ResponseDataV3{
		Status:  cg.Success,
		Message: message,
	}

	utils.WriteResponse(res, respData, http.StatusOK)
	return
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

	if len(data) == 0 {
		respData = &utils.ResponseDataV3{
			Status: cg.Fail,
			Message: map[string]string{
				"en": "Not Found",
				"id": "Tidak Ditemukan",
			},
		}
		utils.WriteResponse(res, respData, http.StatusNotFound)
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
