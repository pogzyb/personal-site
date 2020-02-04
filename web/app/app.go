package app

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

type (
	Application struct {
		Config *Config
		Router *gin.Engine
		Templates *template.Template
	}

	//article struct {
	//	Uid		string `json:"uid"`
	//	Title	string `json:"title"`
	//	Label	string `json:"label"`
	//	DateStr	string `json:"date-string"`
	//	Preview	template.HTML `json:"preview"`
	//	Body	template.HTML `json:"body"`
	//}
	//
	//page struct {
	//	LoggedIn	bool
	//	CurrentArt	*article
	//	Articles 	[]*article
	//}
)

func New() (app *Application, err error) {
	app = &Application{}
	app.Config, err = InitConfig()
	app.Router, err = InitRouter()
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (app *Application) Start() {
	app.Router.Run()
}
