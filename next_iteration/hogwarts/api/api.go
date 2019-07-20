// Package api is the entry point of 'magic' service
//
package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
	"path"
	"runtime"
)

// function ServeSwaggerUI allows to display
// API documentation on localhost:9090/
// It find current working dir (cwd)
// And joins it to swagger-ui dir relative path
// Then it loads this dir (and trims it from serve path)
func ServeSwaggerUI(rt *mux.Router) {
	_, cwd, _, _ := runtime.Caller(1)

	ui := path.Join(path.Dir(cwd), "/swaggerui/")

	rt.PathPrefix("/swaggerui/").Handler(
		http.StripPrefix("/swaggerui/", http.FileServer(http.Dir(ui))))
}

// InitHogwarts starts the Hogwarts service
func InitHogwarts(db *sql.DB) (err error) {
	rt := mux.NewRouter().StrictSlash(true) // handle trailing slash on each route

	type W = http.ResponseWriter
	type R = http.Request

	// Swagger handling
	ServeSwaggerUI(rt)

	// GET actions
	rt.Methods("GET").Path("/").HandlerFunc(func(w W, r *R) { err = Index(&w, r) })

	// POST actions
	rt.Methods("POST").Path("/hogwarts/protect").HandlerFunc(func(w W, r *R) { err = ProtectHogwarts(&w, r, db) })

	// /students/attend

	http.Handle("/", rt)

	return err
}
