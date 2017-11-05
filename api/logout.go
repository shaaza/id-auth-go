package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"micro-auth/domain"
	"micro-auth/serializer"
	"micro-auth/service"
	"net/http"
)

type LogoutHandler struct {
	UserService service.UserService
}

func (h LogoutHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	sessionID := vars["session_id"]
	if len(sessionID) == 0 {
		e := fmt.Errorf("required param session id not provided")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(domain.ErrToJSON(e, http.StatusBadRequest))
		return
	}

	logoutErr := h.UserService.Logout(sessionID)
	if logoutErr != nil {
		e := fmt.Errorf("user logout failed: %s", logoutErr.Error())
		rw.WriteHeader(logoutErr.Code())
		rw.Write(domain.ErrToJSON(e, logoutErr.Code()))
		return
	}

	successResponse := serializer.LogoutResponse{
		Status: "SUCCESS",
	}

	responseJson, err := json.Marshal(successResponse)
	if err != nil {
		e := fmt.Errorf("could not parse resp json: %s", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(domain.ErrToJSON(e, http.StatusInternalServerError))
	}

	rw.Write(responseJson)
}
