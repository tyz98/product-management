package models

// Product struct represents the Product table in the database
type Product struct {
	ID          uint    `json:"id" form:"id" gorm:"primaryKey"`
	Name        string  `json:"name" form:"name"`
	Type        string  `json:"type" form:"type"`
	Price       float64 `json:"price" form:"price"`
	Description string  `json:"description" form:"description"`
	ImagePath   string  `json:"image_path" form:"image_path"`
}
