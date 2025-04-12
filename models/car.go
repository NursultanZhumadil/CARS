package models

type Car struct {
	ID       uint `gorm:"primaryKey"`
	Brand    string
	CarModel string
	Year     int
	Price    float64
	Mileage  int
}
