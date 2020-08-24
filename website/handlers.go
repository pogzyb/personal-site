package website

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/julienschmidt/httprouter"
)

// simple data storage for articles
var ArticleStorage Articles

type (
	Article struct {
		Title			string `json:"title"`
		Link			string `json:"link"`
		TmplName		string `json:"template"`
		HTMLTmplName	string `json:"html_template"`
		Date			string `json:"date"`
		FrontPage		bool   `json:"front_page"`
	}

	Articles struct {
		AllArticles []Article `json:"articles"`
	}
)

// read article data from json file before starting the app
func init()  {
	cwd, _ := os.Getwd()
	pathToJSON := path.Join(cwd, "data/articles.json")
	f, err := os.Open(pathToJSON)
	if err != nil {
		log.Fatalf("Could not open json file [%s]\n%s", pathToJSON, err)
	}
	defer f.Close()
	j, _ := ioutil.ReadAll(f)
	err = json.Unmarshal(j, &ArticleStorage); if err != nil {
		log.Fatalf("Could not unmarhsall json!\n%s", err)
	}
}

func getArticleByLink(link string) (Article, error) {
	for _, article := range ArticleStorage.AllArticles {
		if link == article.Link {
			return article, nil
		}
	}
	return Article{}, errors.New(fmt.Sprintf("No article for link: %s", link))
}

func (s Site) index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	index := s.Templates.Lookup("index.gohtml")
	if err := index.ExecuteTemplate(w, "index", ArticleStorage); err != nil {
		log.Fatalf("Could not write response for 'index' page... %s", err.Error())
	}
}

func (s Site) resume(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resume := s.Templates.Lookup("resume.gohtml")
	if err := resume.ExecuteTemplate(w, "resume", nil); err != nil {
		log.Fatalf("Could not write response for 'resume' page... %s", err.Error())
	}
}

func (s Site) projects(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resume := s.Templates.Lookup("projects.gohtml")
	if err := resume.ExecuteTemplate(w, "projects", nil); err != nil {
		log.Fatalf("Could not write response for 'projects' page... %s", err.Error())
	}
}

func (s Site) articles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resume := s.Templates.Lookup("articles.gohtml")
	if err := resume.ExecuteTemplate(w, "articles", ArticleStorage); err != nil {
		log.Fatalf("Could not write response for 'resume' page... %s", err.Error())
	}
}

func (s Site) article(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	link := ps.ByName("link")
	a, err := getArticleByLink(link)
	if err != nil {
		log.Fatal(err)
	}
	article := s.Templates.Lookup(a.TmplName)
	if err := article.ExecuteTemplate(w, a.TmplName, nil); err != nil {
		log.Fatalf("Could not write response for 'resume' page... %s", err.Error())
	}
}