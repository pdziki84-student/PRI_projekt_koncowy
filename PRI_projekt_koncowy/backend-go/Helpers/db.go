package Helpers

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func init() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://pd:abc1234@cluster0.r2ntk.mongodb.net/ToDos?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	if err == nil {
		log.Println("Connected to DB!")
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	DB = client.Database("ToDos")
}
