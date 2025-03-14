package router

import (
	"net/http"
	"nimblestack/database"
	"nimblestack/router/apis"
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
	// serving auth page
	r.mux.HandleFunc("/auth", handlers.Auth)
	//serving protected pages
	r.mux.HandleFunc("/student/dashboard", middleware.CheckAuth(handlers.StudentDashHandler, r.jwtSercet))
	r.mux.HandleFunc("/supervisor/dashboard", middleware.CheckAuth(handlers.SupervisorDashHandler, r.jwtSercet))
	r.mux.HandleFunc("/coordinator/dashboard", middleware.CheckAuth(handlers.CoordinatorDashHandler, r.jwtSercet))
	// apis section
	authHander := apis.NewAuthApi(r.queries, r.jwtSercet)
	r.mux.HandleFunc("/api/login", authHander.Login)
	r.mux.HandleFunc("/api/register", authHander.Reqister)
	// serving protected apis
	userHandler := apis.NewUserApi(r.queries)
	r.mux.HandleFunc("/api/me", middleware.CheckAuth(userHandler.GetCurrentUSer, r.jwtSercet))
}

func (r *Router) Handler() http.Handler {
	return r.mux
}
