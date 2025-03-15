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

type safeSupervisor struct {
	Id        int64
	Firstname string
	Lastname  string
	Email     string
}
type safeStudent struct {
	Id         int64
	Firstname  string
	Lastname   string
	Email      string
	Supervisor safeSupervisor
}

func (h *DashApi) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	var safeStudents []safeStudent
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	students, err := h.Queries.GetAllStudents(ctx)
	if err != nil {
		http.Error(w, "Error with getting students", http.StatusInternalServerError)
		log.Println("Error with getting students: ", err)
		return
	}
	for _, student := range students {
		if student.Supervisorid.Int64 != 0 {
			supervisor, err := h.Queries.GetSupervisorById(ctx, student.Supervisorid.Int64)
			if err != nil {
				http.Error(w, "Error with getting student supervisor", http.StatusInternalServerError)
				log.Println("Error with getting students: ", err)
				return
			}
			safeStudent := safeStudent{
				Id:        student.Studentid,
				Firstname: student.Firstname,
				Lastname:  student.Lastname,
				Email:     student.Email,
				Supervisor: safeSupervisor{
					Id:        supervisor.Supervisorid,
					Firstname: supervisor.Firstname,
					Lastname:  supervisor.Lastname,
					Email:     supervisor.Email,
				},
			}
			safeStudents = append(safeStudents, safeStudent)
		} else {
			safeStudent := safeStudent{
				Id:         student.Studentid,
				Firstname:  student.Firstname,
				Lastname:   student.Lastname,
				Email:      student.Email,
				Supervisor: safeSupervisor{},
			}
			safeStudents = append(safeStudents, safeStudent)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(safeStudents); err != nil {
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
		log.Printf("Failed to encode response: %v", err)
	}
}
