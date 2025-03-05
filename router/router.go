package router

import (
	"net/http"
	"nimblestack/database"
	"nimblestack/router/handlers"
	"nimblestack/router/middleware"
)

type Router struct {
	queries   *database.Queries
	jwtSercet []byte
	mux       *http.ServeMux
}

func NewRouter(queries *database.Queries, jwtSercet []byte) *Router {
	r := &Router{
		queries:   queries,
		jwtSercet: jwtSercet,
		mux:       http.NewServeMux(),
	}
	r.setupRoutes()
	return r
}

func (r *Router) setupRoutes() {
	userHandler := handlers.NewUserHandler(r.queries, r.jwtSercet)
	dashHandler := handlers.NewDashboardHandler(r.queries)

	r.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	r.mux.HandleFunc("/", handlers.IndexHandler)
	r.mux.HandleFunc("/auth/login", userHandler.Login)
	r.mux.HandleFunc("/auth/register", userHandler.RegisterUser)
	r.mux.HandleFunc("/dashboard", middleware.CheckAuth(dashHandler.DashHandler, r.jwtSercet))
	r.mux.HandleFunc("/admin", middleware.CheckAuth(dashHandler.AuthDash, r.jwtSercet))
	r.mux.HandleFunc("/api/users/{id}/delete", dashHandler.DeleteUser)
	r.mux.HandleFunc("/api/users/{id}/update", dashHandler.UpdateUser)
}

func (r *Router) Handler() http.Handler {
	return r.mux
}
