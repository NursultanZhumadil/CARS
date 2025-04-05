package models

type Car struct {
	ID      int     `json:"id" gorm:"primaryKey"`
	Brand   string  `json:"brand"`
	Model   string  `json:"model"`
	Year    int     `json:"year"`
	Price   float64 `json:"price"`
	Mileage float64 `json:"mileage"`
}
