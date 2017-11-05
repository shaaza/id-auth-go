package api

import (
	"encoding/json"
	"fmt"
	"micro-auth/domain"
	"micro-auth/serializer"
	"micro-auth/service"
	"net/http"
)

type LoginHandler struct {
	UserService service.UserService
}

func (h LoginHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(req.Body)
	var reqBody *serializer.LoginRequest
	err := decoder.Decode(&reqBody)
	if err != nil {
		e := fmt.Errorf("json decode failed: %s", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(domain.ErrToJSON(e, http.StatusBadRequest))
		return
	}

	user, session, loginErr := h.UserService.Login(reqBody)
	if loginErr != nil {
		e := fmt.Errorf("user login failed: %s", loginErr.Error())
		rw.WriteHeader(loginErr.Code())
		rw.Write(domain.ErrToJSON(e, loginErr.Code()))
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
		e := fmt.Errorf("could not parse resp json: %s", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(domain.ErrToJSON(e, http.StatusInternalServerError))
	}

	rw.Write(responseJson)
}
