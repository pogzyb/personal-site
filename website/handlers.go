package website

import (
	//"errors"
	//"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	Page struct {
		Title		string
		Content 	[]interface{}
	}

	Articles []Article
)

func (s Site) index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	indexTmpl := s.Templates.Lookup("index.gohtml")
	if err := indexTmpl.ExecuteTemplate(w, "index", nil); err != nil {
		log.Fatalf("Could not write response for 'index' page: %v", err)
	}
}

func (s Site) resume(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resumeTmpl := s.Templates.Lookup("resume.gohtml")
	if err := resumeTmpl.ExecuteTemplate(w, "resume", nil); err != nil {
		log.Fatalf("could not write response for 'resume' page: %v", err)
	}
}

func (s Site) projects(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	projectsTmpl := s.Templates.Lookup("projects.gohtml")
	if err := projectsTmpl.ExecuteTemplate(w, "projects", nil); err != nil {
		log.Fatalf("could not write response for 'projects' page: %v", err)
	}
}

func (s Site) articles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var many Articles
	queryResult := database.Find(&many)
	log.Println(queryResult.RowsAffected)
	if queryResult.Error != nil {
		log.Printf("could not query articles: %v", queryResult.Error)
	}
	articlesTmpl := s.Templates.Lookup("articles.gohtml")
	if err := articlesTmpl.ExecuteTemplate(w, "articles", many); err != nil {
		log.Fatalf("could not write response for 'articles' page: %v", err)
	}
}

//func (s Site) article(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	link := ps.ByName("link")
//	a, err := getArticleByLink(link)
//	if err != nil {
//		log.Fatal(err)
//	}
//	article := s.Templates.Lookup(a.TmplName)
//	if err := article.ExecuteTemplate(w, a.TmplName, nil); err != nil {
//		log.Fatalf("Could not write response for 'resume' page... %s", err.Error())
//	}
//}