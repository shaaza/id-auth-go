package api

import (
	"encoding/json"
	"fmt"
	"micro-auth/serializer"
	"micro-auth/service"
	"net/http"
)

type LoginHandler struct {
	UserService service.UserService
}

func (h LoginHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var reqBody *serializer.LoginRequest
	err := decoder.Decode(&reqBody)
	if err != nil {
		http.Error(rw, fmt.Sprintf("json decode failed: %s", err.Error()), http.StatusBadRequest)
		return
	}

	user, session, loginErr := h.UserService.Login(reqBody)
	if loginErr != nil {
		http.Error(rw, fmt.Sprintf("user login failed: %s", loginErr.Error()), loginErr.Code())
		return
	}

	successResponse := serializer.LoginResponse{
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		SessionId:   session.Id,
	}

	responseJson, err := json.Marshal(successResponse)
	if err != nil {
		http.Error(rw, fmt.Sprintf("could not parse resp json: %s", err.Error()), http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(responseJson)
}
