package main

import (
	"fmt"
	"net/http"

	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"google.golang.org/appengine"
)

type Customer struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	router := httprouter.New()
	c := cors.New(cors.Options{
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "X-Requested-With"},
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "OPTIONS", "HEADS"},
		AllowCredentials: false,
		MaxAge:           1200,
	})
	router.OPTIONS("/", options)
	router.GET("/", Index)
	router.POST("/", postTest)
	router.GET("/hello/:name", Hello)
	http.Handle("/", c.Handler(router))
	appengine.Main()
}

func options(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(200)
}
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var std = Customer{Name: ps.ByName("name"), Age: 10}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(std)
}

func postTest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var customer Customer
	if r.Body == nil {
		http.Error(w, " content is blank ", 400)
	}
	json.NewDecoder(r.Body).Decode(&customer)
	w.Header().Set("Content-Type", "application/json")
	customer.Name = " Hello ! " + customer.Name
	json.NewEncoder(w).Encode(customer)

}
