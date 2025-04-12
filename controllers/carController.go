package controllers

import (
	"net/http"

	"awesomeProject12/database"
	"awesomeProject12/models"
	"github.com/gin-gonic/gin"
)

func GetCars(c *gin.Context) {
	var cars []models.Car
	if err := database.DB.Find(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Көліктерді жүктеу сәтсіз"})
		return
	}
	c.JSON(http.StatusOK, cars)
}

func GetCarByID(c *gin.Context) {
	id := c.Param("id")
	var car models.Car
	if err := database.DB.First(&car, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Көлік табылмады"})
		return
	}
	c.JSON(http.StatusOK, car)
}

func CreateCar(c *gin.Context) {
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Деректер дұрыс емес"})
		return
	}
	if err := database.DB.Create(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Көлікті қосу сәтсіз"})
		return
	}
	c.JSON(http.StatusCreated, car)
}

func UpdateCar(c *gin.Context) {
	id := c.Param("id")
	var car models.Car
	if err := database.DB.First(&car, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Көлік табылмады"})
		return
	}
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Деректер дұрыс емес"})
		return
	}
	database.DB.Save(&car)
	c.JSON(http.StatusOK, car)
}

func DeleteCar(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Car{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Көлікті жою сәтсіз"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Көлік жойылды"})
}
