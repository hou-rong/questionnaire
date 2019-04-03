package models

type Category struct {
	ID int `gorm:"primary_key" json:"category_id"`
	Name string `json:"category_name"`
}
