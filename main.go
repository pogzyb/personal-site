package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"pogzyb/website"
)

func main() {
	site := website.NewSite()
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Printf("Starting server on %s", port)
	log.Fatal(http.ListenAndServe(port, site.Router))
}
