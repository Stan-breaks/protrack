package handlers

import (
	"log"
	"net/http"
	"nimblestack/views"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if err := views.Index().Render(r.Context(), w); err != nil {
		log.Println("Error rendering view:", err)
	}
}
