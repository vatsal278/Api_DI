package controller

import "net/http"

type IController interface {
	AllEmployee() func(w http.ResponseWriter, r *http.Request)
	InsertEmployee() func(w http.ResponseWriter, r *http.Request)
	UpdateEmployee() func(w http.ResponseWriter, r *http.Request)
	DeleteEmployee() func(w http.ResponseWriter, r *http.Request)
}
