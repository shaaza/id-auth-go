package api

import (
	"encoding/json"
	"fmt"
	"micro-auth/domain"
	"micro-auth/serializer"
	"micro-auth/service"
	"net/http"
)

type SignupHandler struct {
	UserService service.UserService
}

func (h SignupHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(req.Body)
	var reqBody *serializer.SignupRequest
	err := decoder.Decode(&reqBody)
	if err != nil {
		e := fmt.Errorf("json decode failed: %s", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(domain.ErrToJSON(e, http.StatusBadRequest))
		return
	}

	validateErr := h.validateParams(reqBody)
	if validateErr != nil {
		rw.WriteHeader(validateErr.Code())
		rw.Write(domain.ErrToJSON(validateErr, validateErr.Code()))
	}

	registerErr := h.UserService.Register(reqBody)
	if registerErr != nil {
		e := fmt.Errorf("user registration failed: %s", registerErr.Error())
		rw.WriteHeader(registerErr.Code())
		rw.Write(domain.ErrToJSON(e, registerErr.Code()))
		return
	}

	successResponse := serializer.SignupResponse{
		Status: "SUCCESS",
	}

	response, err := json.Marshal(successResponse)
	if err != nil {
		e := fmt.Errorf("json marshal failed for response: %s", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(domain.ErrToJSON(e, http.StatusInternalServerError))
	}

	rw.Write(response)
}

func (h SignupHandler) validateParams(reqBody *serializer.SignupRequest) *domain.Error {
	if reqBody.Username == "" || reqBody.Password == "" {
		return domain.NewError("required params not provided", http.StatusBadRequest)
	}

	return nil
}
