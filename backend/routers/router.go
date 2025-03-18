package routers

import (
	"net/http"
	"tokenSystem/controllers"

	"github.com/rs/cors"
)

func Router() http.Handler {
	mux := http.NewServeMux()

	// Define routes
	mux.HandleFunc("/student", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateStudnet(w, r, controllers.Client)
	})
	mux.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetAllstudents(w, r, controllers.Client)
	})
	mux.HandleFunc("/login", controllers.AuthenticateLDAP)

	// Apply CORS middleware
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(mux)
}
