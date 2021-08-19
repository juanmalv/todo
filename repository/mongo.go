package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Client Initializes mongo client for global usage
var Client *mongo.Client
var err error

func init() {
	// Set client options for local testing
	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	//Set client options for cloud testing
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://juanmalv:Juanm4lv!@cluster0.upof5.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	Client, err = mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// Check the connection
	err = Client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}
