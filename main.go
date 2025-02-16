package main

import (
	"log"
	"net/http"
	"nimblestack/views"

	"github.com/a-h/templ"
)

func main() {
	mux := http.NewServeMux()

	// Serve static files from the "public" directory
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))

	// Handle the root route

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Render the Index template from the views package.
		if err := views.Index().Render(r.Context(), w); err != nil {
			log.Println("Error rendering view:", err)
		}
	})

	log.Println("NimbleStack server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
