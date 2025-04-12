package routes

import (
	"awesomeProject12/controllers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := gin.Default()

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/cars", getCars)
		authorized.POST("/car", createCar)
		authorized.PUT("/car/:id", updateCar)
		authorized.DELETE("/car/:id", deleteCar)
	}

	return router
}
