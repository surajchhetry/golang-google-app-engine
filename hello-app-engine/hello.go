package main

import (
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
	router.POST("/post", postTest)
	router.GET("/hello/:name", Hello)
	router.GET("/error", displayError)
	http.Handle("/", c.Handler(router))
	appengine.Main()
}

func options(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	OkWithMessage(w, "Loving it ...")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var std = Customer{Name: ps.ByName("name"), Age: 10}
	OkWithData(w, std)
}

func postTest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var customer Customer
	if r.Body == nil {
		Error(w)
	}
	json.NewDecoder(r.Body).Decode(&customer)
	customer.Name = " Hello ! " + customer.Name
	OkWithData(w, customer)

}

func displayError(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Error(w)
}
