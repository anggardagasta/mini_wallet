package delivery

import (
	"encoding/json"
	"github.com/anggardagasta/mini_wallet/models"
	"github.com/anggardagasta/mini_wallet/service/repository/response"
	"github.com/asaskevich/govalidator"
	"net/http"
)

func (hd handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var request models.FormRegister

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		response.ResultError(w, http.StatusInternalServerError, response.MessageInternalError, err)
		return
	}

	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		response.ResultError(w, http.StatusBadRequest, response.MessageBadRequest, err)
		return
	}

	data, err := hd.usersUseCase.RegisterUser(request)
	if err != nil {
		response.ResultError(w, http.StatusBadRequest, response.MessageBadRequest, err)
		return
	}
	response.ResultWithData(w, data, response.MessageSucceed, http.StatusOK)
}

func (hd handler) Auth(w http.ResponseWriter, r *http.Request) {
	var request models.FormAuth

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		response.ResultError(w, http.StatusInternalServerError, response.MessageInternalError, err)
		return
	}

	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		response.ResultError(w, http.StatusBadRequest, response.MessageBadRequest, err)
		return
	}

	data, err := hd.usersUseCase.Auth(request)
	if err != nil {
		response.ResultError(w, http.StatusBadRequest, response.MessageBadRequest, err)
		return
	}
	response.ResultWithData(w, data, response.MessageSucceed, http.StatusOK)
}

