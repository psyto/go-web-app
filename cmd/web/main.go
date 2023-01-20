package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/psyto/go-web-app/pkg/config"
	"github.com/psyto/go-web-app/pkg/handlers"
	"github.com/psyto/go-web-app/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
