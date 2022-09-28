package handlers

import (
	"net/http"

	"github.com/MrBomber0x001/reservation-system/pkg/render"
)

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.template.htm")
}

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.template.htm")
}
