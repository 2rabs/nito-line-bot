package controller

import "net/http"

type Router struct {
	http.Handler
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Handle(route string, handler http.Handler) {
	http.Handle(route, handler)
}

func (r *Router) HandleFunc(route string, handler http.HandlerFunc) {
	http.HandleFunc(route, handler)
}
