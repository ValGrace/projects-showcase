package routes

import (
	"net/http"

	"github.com/ValGrace/projects-showcase/src/controllers"
	"github.com/ValGrace/projects-showcase/src/utils"
	"github.com/gorilla/mux"
)

var RegisterProjectRoutes = func(router *mux.Router) {
	router.HandleFunc("/projects/", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/projects", controllers.GetProject).Methods("GET")
	router.HandleFunc("/projects/{projectId}", controllers.GetProjectById).Methods("GET")
	router.HandleFunc("/projects/{projectId}", controllers.UpdateProject).Methods("PUT")
	router.HandleFunc("/projects/{projectId}", controllers.DeleteProject).Methods("DELETE")
	router.HandleFunc("/users/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", controllers.GetUser).Methods("GET")
	router.Handle("/users/{userId}", utils.AuthMiddleware(http.HandlerFunc(controllers.GetUserById))).Methods("GET")
	router.HandleFunc("/users/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{userId}", controllers.DeleteUser).Methods("DELETE")

}
