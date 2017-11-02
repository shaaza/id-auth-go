package api

import (
	"net/http"

	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"micro-auth/config"
	"micro-auth/db"
	"micro-auth/service"
)

func NewUserService(db db.DB) service.UserService {
	return service.UserServiceImpl{
		Database: db,
	}
}

func NewSignupHandler(db db.DB) SignupHandler {
	return SignupHandler{
		UserService: NewUserService(db),
	}
}

func Server(db db.DB) *negroni.Negroni {
	router := mux.NewRouter()
	router.HandleFunc("/ping", http.HandlerFunc(healthCheck)).Methods("GET")
	router.Handle("/users", NewSignupHandler(db)).Methods("POST")

	recovery := negroni.NewRecovery()
	recovery.PrintStack = false

	server := negroni.New(recovery)
	server.UseHandler(router)
	return server

}

func StartServer(server *negroni.Negroni, config config.AppServer) {
	server.Run(fmt.Sprintf(":%d", config.Port))
}
