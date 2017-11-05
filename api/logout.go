package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"micro-auth/serializer"
	"micro-auth/service"
	"net/http"
)

type LogoutHandler struct {
	UserService service.UserService
}

func (h LogoutHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	sessionID := vars["session_id"]
	if len(sessionID) == 0 {
		http.Error(rw, "required param session id not provided: %s", http.StatusBadRequest)
		return
	}

	logoutErr := h.UserService.Logout(sessionID)
	if logoutErr != nil {
		http.Error(rw, fmt.Sprintf("user logout failed: %s", logoutErr.Error()), logoutErr.Code())
		return
	}

	successResponse := serializer.LogoutResponse{
		Status: "SUCCESS",
	}

	responseJson, err := json.Marshal(successResponse)
	if err != nil {
		http.Error(rw, fmt.Sprintf("could not parse resp json: %s", err.Error()), http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(responseJson)
}
