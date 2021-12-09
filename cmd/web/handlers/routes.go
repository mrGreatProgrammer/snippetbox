package handlers

import (
	"net/http"

	// "github.com/mrGreatProgrammer/snippetbox/cmd/web/handlers"
)

func (app *Application) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet", app.ShowSnippet)
	mux.HandleFunc("/snippet/create", app.CreateSnippet)

	fileServer := http.FileServer(http.Dir("ui/static/"))
	
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}