package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"google.golang.org/appengine"

	"github.com/surajchhetry/golang-google-app-engine/rest"
	"google.golang.org/appengine/datastore"
	"encoding/json"
	"google.golang.org/appengine/log"
	"strconv"
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
	router.GET("/", index)
	router.POST("/", postTest)
	router.GET("/hello/:name", Hello)
	//router.GET("/error", displayError)
	http.Handle("/", c.Handler(router))
	appengine.Main()
}

func options(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}
func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := appengine.NewContext(r)
	q := datastore.NewQuery("customers").Order("Name")
	var customers []Customer
	_, err := q.GetAll(ctx, &customers)
	if err != nil {
		//log.Errorf(ctx, "fetching people: %v", err)
		rest.ErrorWithMessage(w,err.Error())
		return
	}
	rest.OkWithData(w, customers)

}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a, _  := strconv.Atoi( ps.ByName("name") )
	q := datastore.NewQuery("customers").Filter("Age >",a)
	ctx := appengine.NewContext(r)
	log.Infof(ctx," Entered Value :: " +  strconv.Itoa(a))
	var customers []Customer
	_, err := q.GetAll(ctx, &customers)
	if err != nil {
		log.Errorf(ctx, "fetching people: %v", err)
		return
	}

	rest.OkWithData(w, customers)
}

func postTest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var customer Customer
	if r.Body == nil {
		rest.Error(w)
	}
	json.NewDecoder(r.Body).Decode(&customer)
	// Store Data
	ctx := appengine.NewContext(r)
	_, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "customers", nil), &customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rest.Ok(w)
}

func displayError(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rest.Error(w)
}
