package app

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func InitRouter() *httprouter.Router {
	router := httprouter.New()
	return router
}

func (app *Application) indexPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	index := app.Templates.Lookup("index.html")
	if err := index.ExecuteTemplate(w, "index", page{}); err != nil {
		log.Fatalf("Could not write response for 'index' page... %s", err.Error())
	}
}
