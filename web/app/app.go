package app

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"os"
)

type (
	Application struct {
		Config    *Config
		Router    *httprouter.Router
		Templates     *template.Template
	}

	article struct {
		Uid		string `json:"uid"`
		Title	string `json:"title"`
		Label	string `json:"label"`
		DateStr	string `json:"date-string"`
		Preview	template.HTML `json:"preview"`
		Body	template.HTML `json:"body"`
	}

	page struct {
		LoggedIn	bool
		CurrentArt	*article
		Articles 	[]*article
	}
)

func New() (app *Application, err error) {
	app = &Application{}

	app.Templates, err = InitTemplates()
	app.Config, err = InitConfig()
	app.Router = InitRouter()

	app.Router.GET("/", app.indexPage)

	if err != nil {
		return nil, err
	}
	return app, nil
}

func (app *Application) Start() {
	_ = http.ListenAndServe(os.Getenv("APP_PORT"), app.Router)
}
