package routes

import (
	"awesomeProject12/controllers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Машиналарға арналған маршруттар
	router.HandleFunc("/cars", controllers.GetCars).Methods("GET")
	router.HandleFunc("/car/{id}", controllers.GetCar).Methods("GET")
	router.HandleFunc("/car", controllers.CreateCar).Methods("POST")
	router.HandleFunc("/car/{id}", controllers.UpdateCar).Methods("PUT")
	router.HandleFunc("/car/{id}", controllers.DeleteCar).Methods("DELETE")

	return router
}
