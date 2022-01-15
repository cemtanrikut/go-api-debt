package router

import (
	"context"
	"log"
	"net/http"
	"time"

	api "github.com/cemtanrikut/go-api-debt/api/debt"
	apiDebt "github.com/cemtanrikut/go-api-debt/api/debt"
	apiUser "github.com/cemtanrikut/go-api-debt/api/user"

	"github.com/patrickmn/go-cache"

	"github.com/cemtanrikut/go-api-debt/db"
	"github.com/cemtanrikut/go-api-debt/helper"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var ctx context.Context
var userCollection, debtCollection *mongo.Collection
var router *mux.Router
var c *cache.Cache

//User handler
func MuxUserHandler() {
	router, ctx, client, userCollection = db.MongoClient("user_collection")

	router.HandleFunc("/api/user/signup", signUp).Methods(http.MethodPost)
	router.HandleFunc("/api/user/{email}", getUser).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func signUp(w http.ResponseWriter, r *http.Request) {
	result := apiUser.SignUp(w, r, client, userCollection)
	byteRes := helper.JsonMarshal(result)
	w.Write(byteRes)
}
func getUser(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	result := apiUser.GetUser(email, w, r, client, userCollection)
	byteRes := helper.JsonMarshal(result)
	w.Write(byteRes)

	c = cache.New(5*time.Minute, 5*time.Minute)
	c.Set("user_id", result.Data[0], cache.NoExpiration)
}

//Debt Handler
func MuxDebtHandler() {
	router, ctx, client, debtCollection = db.MongoClient("debt_collection")

	router.HandleFunc("/api/debt/add", addDebt).Methods(http.MethodPost)
	router.HandleFunc("/api/debt/update", updateDebt).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func addDebt(w http.ResponseWriter, r *http.Request) {
	userID, found := c.Get("user_id")
	if found {
		result := apiDebt.AddDebt(w, r, client, debtCollection, userID.(string))
		byteRes := helper.JsonMarshal(result)
		w.Write(byteRes)
	}
}
func updateDebt(w http.ResponseWriter, r *http.Request) {
	result := api.UpdateDebt(w, r, debtCollection)
	byteRes := helper.JsonMarshal(result)
	w.Write(byteRes)
}
