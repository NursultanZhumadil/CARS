package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Car struct {
	ID      uint    `json:"id" gorm:"primaryKey"`
	Brand   string  `json:"brand"`
	Model   string  `json:"model"`
	Year    int     `json:"year"`
	Price   float64 `json:"price"`
	Mileage float64 `json:"mileage"`
}

var DB *gorm.DB

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	// PostgreSQL-ға қосылу
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// DB арқылы AutoMigrate
	db.AutoMigrate(&Car{}) // Кесте жасау

	// db-ді глобал айнымалыға сақтау
	DB = db

	fmt.Println("✅ Connected to DB")
}

func main() {
	initDB()

	router := gin.Default()

	router.GET("/cars", getCars)
	router.GET("/car/:id", getCar)
	router.POST("/car", createCar)
	router.PUT("/car/:id", updateCar)
	router.DELETE("/car/:id", deleteCar)

	router.Run(":8080")
}

func getCars(c *gin.Context) {
	var cars []Car
	if err := DB.Find(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cars)
}

func getCar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var car Car
	if err := DB.First(&car, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Car not found"})
		return
	}
	c.JSON(http.StatusOK, car)
}

func createCar(c *gin.Context) {
	var car Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := DB.Create(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, car)
}

func updateCar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var car Car
	if err := DB.First(&car, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Car not found"})
		return
	}

	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	car.ID = uint(id)
	if err := DB.Save(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, car)
}

func deleteCar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DB.Delete(&Car{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted"})
}
