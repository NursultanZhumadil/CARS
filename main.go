package main

import (
	"awesomeProject12/database" // database пакетін өзіңіздің жобаңызға сәйкес орнатыңыз
	"awesomeProject12/models"   // models пакетін өзіңіздің жобаңызға сәйкес орнатыңыз
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Register route
func registerUser(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Register logic here (e.g. save to DB)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login user handler
func loginUser(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Authentication logic here (e.g. verify credentials, issue JWT token)
	c.JSON(http.StatusOK, gin.H{"token": "your.jwt.token.here"})
}

// Get all cars
func getCars(c *gin.Context) {
	var cars []models.Car
	if err := database.DB.Find(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cars"})
		return
	}
	c.JSON(http.StatusOK, cars)
}

// Get a specific car by ID
func getCar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var car models.Car
	result := database.DB.First(&car, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Car not found"})
		return
	}
	c.JSON(http.StatusOK, car)
}

// Create a new car
func createCar(c *gin.Context) {
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create car"})
		return
	}
	c.JSON(http.StatusCreated, car)
}

// Update an existing car
func updateCar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var car models.Car
	if err := database.DB.First(&car, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Car not found"})
		return
	}

	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	car.ID = uint(id)
	if err := database.DB.Save(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update car"})
		return
	}
	c.JSON(http.StatusOK, car)
}

// Delete a car
func deleteCar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := database.DB.Delete(&models.Car{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Car not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted"})
}

func main() {
	database.ConnectDatabase() // <<< ҚОСУ КЕРЕК

	router := gin.Default()

	// Роуттар
	router.POST("/register", registerUser)
	router.POST("/login", loginUser)
	router.GET("/cars", getCars)
	router.GET("/car/:id", getCar)
	router.POST("/car", createCar)
	router.PUT("/car/:id", updateCar)
	router.DELETE("/car/:id", deleteCar)

	router.Run(":8080")
}
