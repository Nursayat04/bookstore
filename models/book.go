package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model `json:"-"`
	ID         uint     `json:"id" gorm:"primarykey"`
	Title      string   `json:"title"`
	AuthorID   uint     `json:"author_id"`
	Author     Author   `json:"-" gorm:"foreignKey:AuthorID"`
	CategoryID uint     `json:"category_id"`
	Category   Category `json:"-" gorm:"foreignKey:CategoryID"`
	Price      float64  `json:"price"`
}
