package routes

import (
	"teachable-test/src/controller"

	"github.com/gorilla/mux"
)

func StartRoutes(router *mux.Router) {
	router.HandleFunc("/courses", controller.GetCourses).Methods("GET")
}
