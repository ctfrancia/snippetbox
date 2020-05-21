package main

import (
	"errors"
	"fmt"

	// "html/template"
	"net/http"
	"strconv"

	"github.com/ctfrancia/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "home.page.tmpl", &templateData{Snippets: s})

	/*
		data := &templateData{Snippets: s}

			files := []string{
				"./ui/html/home.page.tmpl",
				"./ui/html/base.layout.tmpl",
				"./ui/html/footer.partial.tmpl",
			}

			ts, err := template.ParseFiles(files...)
			if err != nil {
				app.serverError(w, err)
				return
			}

			err = ts.Execute(w, data)
			if err != nil {
				app.serverError(w, err)
				http.Error(w, "Internal Server Error", 500)
			}
	*/
}

// ShowSnippet shows a specific snippet
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	app.render(w, r, "show.page.tmpl", &templateData{Snippet: s})
	/*
		files := []string{
			"./ui/html/show.page.tmpl",
			"./ui/html/base.layout.tmpl",
			"./ui/html/footer.partial.tmpl",
		}

		data := &templateData{Snippet: s}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.serverError(w, err)
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.serverError(w, err)
		}
	*/
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// http.MethodPost is a constant "POST"
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := "7"
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("./snippet?id=%id", id), http.StatusSeeOther)
}
