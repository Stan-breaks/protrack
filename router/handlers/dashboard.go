package handlers

import (
	"log"
	"net/http"
	"nimblestack/views"
)

func StudentDashHandler(w http.ResponseWriter, r *http.Request) {
	if err := views.StudentDash().Render(r.Context(), w); err != nil {
		log.Println("Error rendering view:", err)
	}
}
func SupervisorDashHandler(w http.ResponseWriter, r *http.Request) {
	if err := views.SupervisorDashPage().Render(r.Context(), w); err != nil {
		log.Println("Error rendering view:", err)
	}
}

func CoordinatorDashHandler(w http.ResponseWriter, r *http.Request) {
	if err := views.CoordinatorDash().Render(r.Context(), w); err != nil {
		log.Println("Error rendering view:", err)
	}
}
