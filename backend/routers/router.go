package routers

import (
	"net/http"
	"tokenSystem/controllers"
)

func Router() {
	http.HandleFunc("/student", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateStudnet(w, r, controllers.Client)
	})
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetAllstudents(w, r, controllers.Client)
	})
	http.HandleFunc("/login", controllers.AuthenticateLDAP)
}
