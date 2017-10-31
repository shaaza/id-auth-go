package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"micro-auth/service"
)

func NewUserService() service.UserService {
	return service.UserServiceImpl{}
}

func NewSignupHandler() SignupHandler {
	return SignupHandler{
		UserService: NewUserService(),
	}
}

func Server() *negroni.Negroni {
	router := mux.NewRouter()
	router.HandleFunc("/ping", http.HandlerFunc(healthCheck)).Methods("GET")
	router.Handle("/users", NewSignupHandler()).Methods("POST")

	recovery := negroni.NewRecovery()
	recovery.PrintStack = false

	server := negroni.New(recovery)
	server.UseHandler(router)
	return server

}
