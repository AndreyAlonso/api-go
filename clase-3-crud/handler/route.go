package handler

import (
	"github.com/andreyalonso/api-go/clase-3-crud/middleware"
	"net/http"
)

// RoutePerson
func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)

	// handler para POST
	mux.HandleFunc("/v1/persons/create", middleware.Log(middleware.Authentication(h.create)))
	// handler para GET
	mux.HandleFunc("/v1/persons/get-all", middleware.Log(h.getAll))
	// handler para GETByID
	mux.HandleFunc("/v1/persons/get-by-id", middleware.Log(h.getByID))
	// handler para Update
	mux.HandleFunc("/v1/persons/update", middleware.Log(h.update))
	// handler para Delete
	mux.HandleFunc("/v1/persons/delete", middleware.Log(h.delete))
}

// RouteLogin rutas que necesitan un inicio de sesión
func RouteLogin(mux *http.ServeMux, storage Storage) {
	h := newLogin(storage)

	mux.HandleFunc("/v1/login", h.login)
}
