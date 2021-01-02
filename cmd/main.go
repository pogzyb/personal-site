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
	address := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	log.Printf("starting server at %s\n", address)
	log.Fatal(http.ListenAndServe(address, site.Router))
}
