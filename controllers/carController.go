package controllers

import (
	"awesomeProject12/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var cars = []models.Car{
	{ID: 1, Brand: "Toyota", Model: "Camry", Year: 2020, Price: 25000, Mileage: 15000},
	{ID: 2, Brand: "BMW", Model: "X5", Year: 2018, Price: 35000, Mileage: 30000},
}

// Барлық машиналарды алу
func GetCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

// Белгілі бір машинаны алу
func GetCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for _, item := range cars {
		if fmt.Sprintf("%d", item.ID) == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}

// Жаңа машина қосу
func CreateCar(c *gin.Context) {
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, car)
}

// Өнімді жаңарту
func UpdateCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updatedCar models.Car
	_ = json.NewDecoder(r.Body).Decode(&updatedCar)

	for index, item := range cars {
		if fmt.Sprintf("%d", item.ID) == id {
			cars[index] = updatedCar
			updatedCar.ID = item.ID
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedCar)
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}

// Машинаны өшіру
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for index, item := range cars {
		if fmt.Sprintf("%d", item.ID) == id {
			cars = append(cars[:index], cars[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Car deleted successfully")
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}
