package main

import (
    "net/http"
    "github.com/joelewis/guerilla-seed/controllers"
    . "github.com/gorilla/mux"
)

func RegisterRoutes() *Router {
    r := NewRouter()
    s := http.StripPrefix("/static/", http.FileServer(http.Dir(STATIC_DIR)))
    r.PathPrefix("/static/").Handler(s)
    r.HandleFunc("/", controllers.HomeHandler)
    return r
}
