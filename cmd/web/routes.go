package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet?id=1", app.showSnippet)     // ANY
	mux.HandleFunc("/snippet/create", app.createSnippet) // Post

	fileServer := http.FileServer(http.Dir("./ui.static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux

}
