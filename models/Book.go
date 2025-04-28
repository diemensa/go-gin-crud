package models

type Book struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Title  string `json:"title" validate:"required,min=1"`
	Author string `json:"author" validate:"required,min=1"`
	Genre  string `json:"genre" validate:"required,min=1"`
}
