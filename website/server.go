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

var templates *template.Template

type Site struct {
	Router    *httprouter.Router
}

func init() {
	initDatabase()
	initSession()
	initTemplates()
	migrateDatabase()
}

func NewSite() *Site {
	site := Site{Router: initRouter()}
	site.Router.GET("/", site.index)
	site.Router.GET("/resume", site.resume)
	site.Router.GET("/projects", site.projects)
	site.Router.GET("/articles", site.articles)
	site.Router.GET("/article/:articlePath", site.article)
	return &site
}

func initRouter() *httprouter.Router {
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	router.NotFound = http.HandlerFunc(NotFoundHandler)
	router.PanicHandler = InternalServerError
	return router
}

func initTemplates() {
	var templateFiles []string
	templatesDir := os.Getenv("TEMPLATES_DIR")
	// read and parse template files
	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		log.Fatalf("could not read from $TEMPLATES_DIR [%s]: %v", templatesDir, err)
	}
	for _, file := range files {
		templateFiles = append(templateFiles, path.Join(templatesDir, file.Name()))
	}
	templates, err = template.ParseFiles(templateFiles...)
	if err != nil {
		log.Fatalf("could not parse html templates: %v", err)
	}
}
