package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cemtanrikut/go-api-debt/api"
	helper "github.com/cemtanrikut/go-api-debt/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type user struct {
	id          primitive.ObjectID `json:"_id"`
	name        string             `json:"name"`
	phone       string             `json:"phone"`
	email       string             `json:"email"`
	avg_income  float32            `json:"avg_income"`
	create_date time.Time          `json:"create_date"`
	update_date time.Time          `json:"create_date"`
	active      bool               `json:"active"`
}

func SignUp(resp http.ResponseWriter, req *http.Request, client *mongo.Client, collection *mongo.Collection) api.Response {
	resp.Header().Set("Content-Type", "application/json")
	var user user
	json.NewDecoder(req.Body).Decode(&user)
	user.create_date = time.Now()
	user.active = true

	checkEmail := CheckEmail(user.email, client, collection)
	if checkEmail {
		return helper.ReturnResponse(http.StatusUnauthorized, "", "This mail address is already exist.")

	}

	_, insertErr := collection.InsertOne(context.Background(), user)
	if insertErr != nil {
		return helper.ReturnResponse(http.StatusBadRequest, "", insertErr.Error())
	}

	jsonResult, jsonError := json.Marshal(user)
	if jsonError != nil {
		return helper.ReturnResponse(http.StatusInternalServerError, "", jsonError.Error())

	}

	return helper.ReturnResponse(http.StatusOK, string(jsonResult), "")

}

func GetUser(email string, resp http.ResponseWriter, req *http.Request, client *mongo.Client, collection *mongo.Collection) api.Response {
	resp.Header().Set("Content-Type", "application/json")
	var user user

	userData := collection.FindOne(context.Background(), bson.M{"email": email, "active": true})
	err := userData.Decode(&user)

	if err != nil {
		return helper.ReturnResponse(http.StatusNotFound, "", err.Error())
	}

	jsonResult, jsonError := json.Marshal(user)
	if jsonError != nil {
		return helper.ReturnResponse(http.StatusInternalServerError, "", err.Error())
	}

	return helper.ReturnResponse(http.StatusOK, string(jsonResult), "")
}

func CheckEmail(email string, client *mongo.Client, collection *mongo.Collection) bool {
	var dbUser user
	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&dbUser)
	fmt.Println("data - ", err)
	if err == nil {
		fmt.Println(email, " already exist")
		return true
	}
	return false
}
