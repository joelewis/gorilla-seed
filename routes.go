package main

import (
	. "github.com/gorilla/mux"
	"github.com/joelewis/guerilla-seed/controllers"
	"net/http"
)

func RegisterRoutes() *Router {
	r := NewRouter()
	s := http.StripPrefix("/static/", http.FileServer(http.Dir(STATIC_DIR)))
	r.PathPrefix("/static/").Handler(s)
	r.HandleFunc("/", controllers.HomeHandler)
	return r
}
