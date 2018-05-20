package store

import "github.com/jinzhu/gorm"

//Repository ...
type Repository struct{
	Db *gorm.DB
}

func (r Repository)GetProducts() Products {
	products := Products{}
	r.Db.Find(&products)
	return products

}
