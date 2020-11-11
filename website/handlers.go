package website

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

type (
	Articles []Article
	Projects []Project

	Page struct {
		Title		string
		Articles
		Projects
	}
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{ Title: "404 Error" }
	notFoundTmpl := templates.Lookup("404.gohtml")
	if err := notFoundTmpl.ExecuteTemplate(w, "404", page); err != nil {
		log.Fatalf("Could not write response for '404' page: %v", err)
	}
}

func InternalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	subject := fmt.Sprintf("500 Error [%s]", time.Now().Format(time.RFC822))
	serverErr := err.(error)
	stack := string(debug.Stack())
	sendEmail(subject, serverErr, stack)
	page := Page{ Title: "500 Error" }
	internalServerErrorTmpl := templates.Lookup("500.gohtml")
	if err := internalServerErrorTmpl.ExecuteTemplate(w, "500", page); err != nil {
		log.Fatalf("Could not write response for '500' page: %v", err)
	}
}

func (s Site) index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	indexTmpl := templates.Lookup("index.gohtml")
	page := Page{ Title: "Home" }
	if err := indexTmpl.ExecuteTemplate(w, "index", page); err != nil {
		log.Fatalf("Could not write response for 'index' page: %v", err)
	}
}

func (s Site) resume(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resumeTmpl := templates.Lookup("resume.gohtml")
	page := Page{ Title: "Resume" }
	if err := resumeTmpl.ExecuteTemplate(w, "resume", page); err != nil {
		log.Fatalf("could not write response for 'resume' page: %v", err)
	}
}

func (s Site) projects(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var allProjects Projects
	queryResult := database.Find(&allProjects)
	if queryResult.Error != nil {
		log.Printf("could not query projects: %v", queryResult.Error)
	}
	projectsTmpl := templates.Lookup("projects.gohtml")
	page := Page{ Title: "Projects", Projects: allProjects }
	if err := projectsTmpl.ExecuteTemplate(w, "projects", page); err != nil {
		log.Fatalf("could not write response for 'projects' page: %v", err)
	}
}

func (s Site) articles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var allArticles Articles
	queryResult := database.Find(&allArticles)
	if queryResult.Error != nil {
		log.Printf("could not query articles: %v", queryResult.Error)
	}
	articlesTmpl := templates.Lookup("articles.gohtml")
	page := Page{ Title: "Articles", Articles: allArticles }
	if err := articlesTmpl.ExecuteTemplate(w, "articles", page); err != nil {
		log.Fatalf("could not write response for 'articles' page: %v", err)
	}
}

func (s Site) article(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var oneArticle Article
	path := ps.ByName("articlePath")
	queryResult := database.Where("path = ?", path).First(&oneArticle)
	if queryResult.Error != nil {
		log.Printf("could not query article: %s: %v", path, queryResult.Error)
	}
	articleTmpl := templates.Lookup(oneArticle.HTMLTemplate)
	if err := articleTmpl.ExecuteTemplate(w, oneArticle.Template, oneArticle); err != nil {
		log.Fatalf("Could not write response for 'article' %s: %v", oneArticle.Title, err)
	}
}