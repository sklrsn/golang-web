package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index Page",
		"GET",
		"/{username}",
		UserInfoHandler,
	},
}

func createRoutes() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.Handler)
	}
	return router
}

func UserInfoHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Userinfo Handler")
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
}
