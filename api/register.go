package api

import (
	"net/http"
	"encoding/json"
	"micro-auth/serializer"
	"micro-auth/service"
)

type SignupHandler struct {
	UserService service.UserService
}


func (h SignupHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var reqBody *serializer.SignupRequest
	err := decoder.Decode(&reqBody)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("json decode failed: " + err.Error()))
		return
	}

	user, registerErr := h.UserService.Register(reqBody)
	if registerErr != nil || user == "" {
		rw.WriteHeader(registerErr.Code())
		rw.Write([]byte("user registration failed: " + registerErr.Error()))
		return
	}

	successResponse := serializer.SignupResponse{
		Status: "SIGNUP_SUCCESSFUL",
	}

	response, err := json.Marshal(successResponse)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	rw.Write(response)
}
