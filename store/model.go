package store

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Title string `json:"title"`
	Image string `json:"image"`
	Price uint64 `json:"price"`
	Rating uint64 `json:"rating"`
}


type Products []Product
