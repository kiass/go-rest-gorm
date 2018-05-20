package main

import (
	"github.com/gorilla/mux"
	"rest-and-go/store"
	"log"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

var err error

func main() {

	db, err = gorm.Open(
		"postgres",
		"host=localhost user=gear"+
			" dbname=gear sslmode=disable password=gear")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&store.Product{})

	var controller = &store.Controller{Repository: store.Repository{Db: db}}
	router := mux.NewRouter()
	router.Path("/api/products").Methods("GET").Name("Index").HandlerFunc(controller.Index)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func dumpData() {
	db.Create(&store.Product{
		Title:  "Products1",
		Image:  "Image1",
		Price:  10.00,
		Rating: 10})

	db.Create(&store.Product{
		Title:  "Products2",
		Image:  "Image2",
		Price:  20.00,
		Rating: 20})

	db.Create(&store.Product{
		Title:  "Products3",
		Image:  "Image3",
		Price:  30.00,
		Rating: 30})

	db.Create(&store.Product{
		Title:  "Products4",
		Image:  "Image4",
		Price:  40.00,
		Rating: 40})

	db.Create(&store.Product{
		Title:  "Products5",
		Image:  "Image5",
		Price:  50.00,
		Rating: 50})
}
