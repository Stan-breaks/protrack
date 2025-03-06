package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"nimblestack/database"
	"nimblestack/views"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

type DashboardHandler struct {
	Queries *database.Queries
}

func NewDashboardHandler(queries *database.Queries) *DashboardHandler {
	return &DashboardHandler{
		Queries: queries,
	}
}

func (h *DashboardHandler) DashHandler(w http.ResponseWriter, r *http.Request) {
	if err := views.Dash().Render(r.Context(), w); err != nil {
		log.Println("Error rendering view:", err)
	}
}

func (h *DashboardHandler) AuthDash(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	tables, err := h.Queries.GetAllTables(ctx)
	if err != nil {
		log.Println("Error with getting tables: ", err)
	}
	if err := views.AdminDashboard(tables).Render(r.Context(), w); err != nil {
		log.Println("Error rendering view: ", err)
	}
}

func (h *DashboardHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "Failed to get user ID", http.StatusInternalServerError)
	}
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	err = h.Queries.DeleteUser(ctx, userID)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *DashboardHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "Failed to get user ID", http.StatusInternalServerError)
	}
	var user struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	err = h.Queries.UpdateUser(ctx, database.UpdateUserParams{
		ID:       userID,
		Username: user.UserName,
		Password: user.Password,
		// Add other fields as needed
	})
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
