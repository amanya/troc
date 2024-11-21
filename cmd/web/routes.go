package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /troc/view/{id}", app.trocView)
	mux.HandleFunc("GET /troc/create", app.trocCreate)
	mux.HandleFunc("POST /troc/create", app.trocCreatePost)

	return mux
}