package store

import (
	"net/http"
	"encoding/json"
)

type Controller struct {
	Repository Repository
}

func (c *Controller) Index(w http.ResponseWriter, req *http.Request) {
	result := c.Repository.GetProducts()
	data, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
