package routes

import (
	"github.com/afzal/go-course/controller"
	"github.com/gorilla/mux"
)

func CourseRouters(router *mux.Router) {

	router.HandleFunc("/courses", controller.GetAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", controller.GetOneCourse).Methods("GET")
	router.HandleFunc("/course", controller.CreateOneCourse).Methods("POST")
	router.HandleFunc("/course/{id}", controller.UpdateOnecourse).Methods("PUT")
	router.HandleFunc("/course/{id}", controller.DeleteOneCourse).Methods("DELETE")
}
