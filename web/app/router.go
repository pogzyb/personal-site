package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() (*gin.Engine, error) {
	router := gin.Default()
	//router.LoadHTMLGlob("templates/*")
	router.GET("/", indexPage)
	return router, nil
}

func indexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"name": "home"})
}