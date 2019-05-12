package activities

import (
	"encoding/json"
	"fmt"
	"github.com/frayeralex/go-api/db"
	. "github.com/frayeralex/go-api/models"
	"github.com/go-bongo/bongo"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func getCollection() *bongo.Collection {
	return db.Connection.Collection(db.Activities)
}

func GetAll(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "application/json")
	data := make([]Activity, 0)
	fmt.Println(request.URL.Query())
	qType := request.URL.Query().Get("type")
	results := getCollection().Find(bson.M{
		"type": qType,
	})
	activity := &Activity{}

	_meta, err := results.Paginate(10, 1)
	if err != nil {
		fmt.Println(err)
	}

	for results.Next(activity) {
		fmt.Println(activity)
		data = append(data, Activity(*activity))
	}

	responseData := make(map[string]interface{})
	responseData["_meta"] = _meta
	responseData["data"] = data

	json.NewEncoder(writer).Encode(responseData)
}

func GetOne(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	activity := &Activity{}
	err := getCollection().FindById(bson.ObjectIdHex(params["id"]), activity)

	if err != nil {
		if _, ok := err.(*bongo.DocumentNotFoundError); ok {
			fmt.Println("document not found")
			writer.WriteHeader(http.StatusNotFound)
		} else {
			fmt.Println("real error " + err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	_ = json.NewEncoder(writer).Encode(activity)
}

func Create(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "application/json")
	activity := &Activity{}
	_ = json.NewDecoder(request.Body).Decode(&activity)
	_ = getCollection().Save(activity)
	writer.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(writer).Encode(activity)
}

func Update(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	activity := &Activity{}
	err := getCollection().FindById(bson.ObjectIdHex(params["id"]), activity)

	if err != nil {
		if _, ok := err.(*bongo.DocumentNotFoundError); ok {
			activity.SetId(bson.ObjectIdHex(params["id"]))
			_ = json.NewDecoder(request.Body).Decode(&activity)
			_ = getCollection().Save(activity)
			writer.WriteHeader(http.StatusCreated)
			_ = json.NewEncoder(writer).Encode(activity)
		} else {
			fmt.Println("real error " + err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	_ = json.NewDecoder(request.Body).Decode(&activity)
	writer.WriteHeader(http.StatusOK)
	_ = getCollection().Save(activity)
	_ = json.NewEncoder(writer).Encode(activity)
}

func Delete(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	activity := &Activity{}

	err := getCollection().FindById(bson.ObjectIdHex(params["id"]), activity)

	if err != nil {
		if _, ok := err.(*bongo.DocumentNotFoundError); ok {
			fmt.Println("document not found")
			writer.WriteHeader(http.StatusNotFound)
		} else {
			fmt.Println("real error " + err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = getCollection().DeleteDocument(activity)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
