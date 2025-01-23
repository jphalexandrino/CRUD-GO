package mongodb

import (
	"context"
	"github.com/jphalexandrino/CRUD-GO/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//IMPORTANT: I will only use MongoDB for future studies, the main one will continue to be Postgres

var collection *mongo.Collection
var ctx = context.Background()

func InitConnection() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("tasker").Collection("tasks")

	logger.Info("Connection to the database successfully established! - MongolDB")
}
