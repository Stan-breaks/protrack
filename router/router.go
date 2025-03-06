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
	// serving static files
	r.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	// serving the index page
	r.mux.HandleFunc("/", handlers.IndexHandler)
	// serving auth pages
	r.mux.HandleFunc("/login", handlers.Login)
	r.mux.HandleFunc("/register", handlers.Register)
	//serving protected pages
	dashHandler := handlers.NewDashboardHandler(r.queries)
	r.mux.HandleFunc("/dashboard", middleware.CheckAuth(dashHandler.DashHandler, r.jwtSercet))
	r.mux.HandleFunc("/admin", middleware.CheckAuth(dashHandler.AuthDash, r.jwtSercet))
}

func (r *Router) Handler() http.Handler {
	return r.mux
}
