package main

import (
	"errors"
	"fmt"

	"net/http"
	"strconv"

	"troc.amanya/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	trocs, err := app.trocs.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Trocs = trocs

	app.render(w, r, http.StatusOK, "home.tmpl", data)
}

func (app *application) trocView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	troc, err := app.trocs.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
			return
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Troc = troc

	app.render(w, r, http.StatusOK, "view.tmpl", data)
}

func (app *application) trocCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new troc..."))
}

func (app *application) trocCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.trocs.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/troc/view/%d", id), http.StatusSeeOther)
}
