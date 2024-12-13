package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey"` //unique identifier to quickly find(1,2,3,4) | common practice
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`  
}

type Book struct {
	BaseModel
	Title       string  `json:"title" gorm:"type:varchar(255);not null"` //255 chars | not empty
	Description string  `json:"description" gorm:"type:text;not null"`
	Category    string  `json:"category" gorm:"type:varchar(255);not null"`
	Trending    bool    `json:"trending" gorm:"not null"`
	CoverImage  string  `json:"cover_image" gorm:"type:varchar(255);not null"`
	OldPrice    float64 `json:"old_price" gorm:"not null"`
	NewPrice    float64 `json:"new_price" gorm:"not null"`

	//Orders []Order `gorm:"many2many:order_books;"` // many-to-many relationship
}
