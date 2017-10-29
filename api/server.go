package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func Server() *negroni.Negroni {
	router := mux.NewRouter()
	router.HandleFunc("/ping", http.HandlerFunc(healthCheck))

	recovery := negroni.NewRecovery()
	recovery.PrintStack = false

	server := negroni.New(recovery)
	server.UseHandler(router)
	return server

}
