package handlers

import (
	"log"
	"net/http"
	"nimblestack/views"
)

func DashHandler(w http.ResponseWriter, r *http.Request) {
	if err := views.Dash().Render(r.Context(), w); err != nil {
		log.Println("Error rendering view:", err)
	}
}
