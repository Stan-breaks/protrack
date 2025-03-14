package apis

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"nimblestack/database"
	"time"
)

type DashApi struct {
	Queries *database.Queries
}

func NewDashApi(queries *database.Queries) *DashApi {
	return &DashApi{
		Queries: queries,
	}
}

func (h *DashApi) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	students, err := h.Queries.GetAllStudents(ctx)
	if err != nil {
		http.Error(w, "Error with getting students", http.StatusInternalServerError)
		log.Println("Error with getting students: ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(students); err != nil {
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
		log.Printf("Failed to encode response: %v", err)
	}
}
