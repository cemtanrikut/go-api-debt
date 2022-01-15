package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/cemtanrikut/go-api-debt/api"
	"github.com/cemtanrikut/go-api-debt/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type debt struct {
	id          primitive.ObjectID `json:"_id"`
	name        string             `json:"name"`
	user_id     string             `json:"user_id"`
	typeof      string             `json:"typeof"`
	amount      float32            `json:"amount"`
	periodicly  bool               `json:"periodicly"`
	start_date  time.Time          `json:"start_date"`
	end_date    time.Time          `json:"end_date"`
	completed   bool               `json:"completed"`
	active      bool               `json:"active"`
	create_date time.Time          `json:"create_date"`
	update_date time.Time          `json:"update_date"`
}

func AddDebt(resp http.ResponseWriter, req *http.Request, client *mongo.Client, collection *mongo.Collection, userID string) api.Response {
	resp.Header().Set("Content-Type", "application/json")
	var debt debt
	json.NewDecoder(req.Body).Decode(&debt)
	debt.completed = false
	debt.create_date = time.Now()
	debt.active = true
	debt.user_id = userID

	_, err := collection.InsertOne(context.Background(), debt)
	if err != nil {
		return helper.ReturnResponse(http.StatusBadRequest, "", err.Error())
	}

	jsonResult, jsonError := json.Marshal(debt)
	if jsonError != nil {
		return helper.ReturnResponse(http.StatusInternalServerError, "", jsonError.Error())

	}

	return helper.ReturnResponse(http.StatusOK, string(jsonResult), "")

}

func UpdateDebt(resp http.ResponseWriter, req *http.Request, collection *mongo.Collection) api.Response {
	resp.Header().Set("Content-Type", "application/json")
	var debt debt
	json.NewDecoder(req.Body).Decode(&debt)

	updatedData, updateErr := collection.UpdateOne(context.Background(), bson.M{"_id": debt.id, "active": true}, bson.D{{"$set",
		bson.D{
			{"name", debt.name},
			{"typeof", debt.typeof},
			{"amount", debt.amount},
			{"amount", debt.amount},
			{"periodicly", debt.periodicly},
			{"start_date", debt.start_date},
			{"end_date", debt.end_date},
			{"completed", debt.completed},
			{"active", debt.active},
			{"update_date", time.Now()},
		},
	}})
	if updateErr != nil {
		return helper.ReturnResponse(http.StatusInternalServerError, "", updateErr.Error())
	}
	jsonResult, err := json.Marshal(updatedData)
	if err != nil {
		return helper.ReturnResponse(http.StatusInternalServerError, "", err.Error())
	}

	return helper.ReturnResponse(http.StatusOK, string(jsonResult), "")
}
