package router

import (
	"net/http"
	"nimblestack/database"
	"nimblestack/router/handlers"
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
	rootHandler := handlers.RootHandler
	r.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	r.mux.HandleFunc("/", rootHandler)
	r.mux.HandleFunc("/auth/login", userHandler.Login)
	r.mux.HandleFunc("/auth/register", userHandler.RegisterUser)
}

func (r *Router) Handler() http.Handler {
	return r.mux
}
