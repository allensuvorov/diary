package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type user struct {
	ID       int
	Email    string
	Password string
	Username string
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Handler - hello\n")
	now := time.Now()
	fmt.Fprintf(w, "Welcome to your daily diary!\nCurrent date and time: %s\n", now.Format(time.Layout))
	log.Print("Handler - buy\n")
}

func main() {

	http.HandleFunc("/", handler)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func connectToMongoDB() {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://allensuvorov:UwEK8f7e3RUo5d9q@cluster0.sivdk.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	databases, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}
	fmt.Println(databases)
}
