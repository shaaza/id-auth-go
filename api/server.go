package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"micro-auth/config"
	"micro-auth/db"
	"micro-auth/service"
	"net/http"
	"os"
)

func NewUserService(db db.DB) service.UserService {
	return service.UserServiceImpl{
		Database: db,
	}
}

func NewLoginHandler(userService service.UserService) LoginHandler {
	return LoginHandler{
		UserService: userService,
	}
}

func NewLogoutHandler(userService service.UserService) LogoutHandler {
	return LogoutHandler{
		UserService: userService,
	}
}

func NewSignupHandler(userService service.UserService) SignupHandler {
	return SignupHandler{
		UserService: userService,
	}
}

func Server(db db.DB) *negroni.Negroni {
	userService := NewUserService(db)
	log.SetOutput(os.Stdout)

	router := mux.NewRouter()
	router.HandleFunc("/ping", http.HandlerFunc(healthCheck)).Methods("GET")
	router.Handle("/users", NewSignupHandler(userService)).Methods("POST")
	router.Handle("/sessions", NewLoginHandler(userService)).Methods("POST")
	router.Handle("/sessions/{session_id}", NewLogoutHandler(userService)).Methods("PUT")

	recovery := negroni.NewRecovery()
	recovery.PrintStack = false

	server := negroni.New(recovery)
	server.UseHandler(router)
	return server

}

func StartServer(server *negroni.Negroni, config config.AppServer) {
	server.Run(fmt.Sprintf(":%d", config.Port))
}
