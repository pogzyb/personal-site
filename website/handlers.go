package website

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s Site) index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	index := s.Templates.Lookup("index.html")
	if err := index.ExecuteTemplate(w, "index", nil); err != nil {
		log.Fatalf("Could not write response for 'index' page... %s", err.Error())
	}
}

func (s Site) resume(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	index := s.Templates.Lookup("resume.html")
	if err := index.ExecuteTemplate(w, "resume", nil); err != nil {
		log.Fatalf("Could not write response for 'resume' page... %s", err.Error())
	}
}