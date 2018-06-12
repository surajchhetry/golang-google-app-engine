package main

import (
	"fmt"
	"net/http"

	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
)

type Customer struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	http.Handle("/", router)
	appengine.Main()
}
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var std = Customer{Name: ps.ByName("name"), Age: 10}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(std)
}
