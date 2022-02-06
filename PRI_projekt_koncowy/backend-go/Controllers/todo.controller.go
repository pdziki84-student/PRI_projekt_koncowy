package Controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"pri/todos/Helpers"
	"pri/todos/Models"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	tasks := Helpers.DB.Collection("ToDoItems")
	c, _ := tasks.Find(context.TODO(), bson.M{})
	var results []Models.ToDoItem
	for c.Next(context.TODO()) {
		var item Models.ToDoItem
		err := c.Decode(&item)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		results = append(results, item)
	}
	json.NewEncoder(w).Encode(results)

	log.Println("GetTasks done")
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	var item Models.ToDoItem
	json.NewDecoder(r.Body).Decode(&item)
	log.Printf("Created task: %v", item)

	collection := Helpers.DB.Collection("ToDoItems")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := collection.InsertOne(ctx, item)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	json.NewEncoder(w).Encode(result)
}

func DeleteTask(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	response.Header().Add("Access-Control-Allow-Origin", "*")
	response.Header().Add("Access-Control-Allow-Methods", "*")
	response.Header().Add("Access-Control-Allow-Headers", "*")
	response.Header().Add("content-type", "application/json")
	collection := Helpers.DB.Collection("ToDoItems")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	vars := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(vars["id"])
	collection.DeleteOne(ctx, bson.M{"_id": id})
	log.Printf("Deleted task: %v", id)
}
