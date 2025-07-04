package router

import (
	"github.com/gorilla/mux"
	"github.com/hiteshchoudhary/mongodb/controller" // Ensure this path is correct
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/courses", controller.GetAllCourses).Methods("GET")
	router.HandleFunc("/api/course/{id}", controller.GetOneCourse).Methods("GET")
	router.HandleFunc("/api/course", controller.CreateCourse).Methods("POST")
	router.HandleFunc("/api/course/{id}", controller.UpdateCourse).Methods("PUT")
	router.HandleFunc("/api/course/{id}", controller.DeleteCourse).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie/{id}", controller.GetOneMovie).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.UpdateMovie).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/my-movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie/{id}/watched", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMoviesHandler).Methods("DELETE")
	return router
}