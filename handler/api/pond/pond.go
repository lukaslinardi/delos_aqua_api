package pondHandler


import (

	"encoding/json"
	"io/ioutil"
	"net/http"

	cg "github.com/lukaslinardi/delos_aqua_api/domain/constants/general"
	pd "github.com/lukaslinardi/delos_aqua_api/domain/model/pond"

	"github.com/lukaslinardi/delos_aqua_api/domain/model/general"
	"github.com/lukaslinardi/delos_aqua_api/domain/utils"
	"github.com/sirupsen/logrus"

	pondService "github.com/lukaslinardi/delos_aqua_api/service/pond"
)


type PondHandler struct {
	Pond pondService.Pond
	conf general.AppService
	log  *logrus.Logger
}

func NewPondHandler(pond pondService.Pond, conf general.AppService, logger *logrus.Logger) PondHandler {
	return PondHandler{
		Pond: pond,
		conf: conf,
		log:  logger,
	}
}


func (ph PondHandler) GetPonds(res http.ResponseWriter, req *http.Request) {
	respData := &utils.ResponseDataV3{
		Status: cg.Fail,
	}

	data, message, err := ph.Pond.GetPonds(req.Context())
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


func (ph PondHandler) InsertPond(res http.ResponseWriter, req *http.Request) {
	respData := &utils.ResponseDataV3{
		Status: cg.Fail,
	}

	var request pd.InsertPond

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

	messages, err := ph.Pond.InsertPond(req.Context(), request)
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
