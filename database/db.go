package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"awesomeProject12/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=sultan05 dbname=car_db port=5432 sslmode=disable TimeZone=Asia/Almaty"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	err = DB.AutoMigrate(&models.Car{})
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}
}
