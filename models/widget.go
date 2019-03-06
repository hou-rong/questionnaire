package models

type Widget struct {
	ID int `gorm:"primary_key" json:"widget_id"`
	Name string `json:"widget_name"`
}
