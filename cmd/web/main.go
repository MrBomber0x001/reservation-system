package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MrBomber0x001/reservation-system/pkg/config"
	"github.com/MrBomber0x001/reservation-system/pkg/handlers"
	"github.com/MrBomber0x001/reservation-system/pkg/render"
)

const PORT = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	serverError := http.ListenAndServe(PORT, nil)
	fmt.Println(fmt.Sprintf("Starting application on port %s", PORT))
	if serverError != nil {
		fmt.Println(`Error Starting Server`)
	}
}
