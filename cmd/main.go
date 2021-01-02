package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"pogzyb/website"
)

func main() {
	site := website.New()
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Printf("starting server on %s\n", port)
	log.Fatal(http.ListenAndServe(port, site.Router))
}
