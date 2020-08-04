package website

import (
	"github.com/julienschmidt/httprouter"
	"net/http"

	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type Site struct {
	Router    *httprouter.Router
	Templates *template.Template
}

func NewSite() *Site {
	router := initRouter()
	templates := initTemplates()
	site := Site{
		Router:    router,
		Templates: templates,
	}
	site.Router.GET("/", site.index)
	site.Router.GET("/resume", site.resume)
	return &site
}

func initRouter() *httprouter.Router {
	router := httprouter.New()
	// add handlers
	//router.GET("/", Site.index)
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	return router
}

func initTemplates() *template.Template {
	var templateFiles []string
	templatesDir := os.Getenv("TEMPLATES_DIR")
	// read and parse template files
	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		log.Fatalf("Could not read from $TEMPLATES_DIR [%s]...\n%s", templatesDir, err.Error())
	}
	for _, file := range files {
		templateFiles = append(templateFiles, path.Join(templatesDir, file.Name()))
	}
	templates, err := template.ParseFiles(templateFiles...)
	if err != nil {
		log.Fatalf("Could not parse html templates...\n%s", err.Error())
	}

	return templates
}
