package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"net/http"

	"github.com/gin-gonic/gin"

	"encoding/json"
	"os"
)

type test struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var testArr = []test{
	{ID: "1", Title: "Blue Train"},
	{ID: "2", Title: "Jeru"},
}

type Secrets struct {
	ConnectionURI string `json:"connectionURI"`
}

func main() {

	// grabbing the connection URI for database from secrets file
	// I can help you set this up on your machine
	file, err := os.Open("secrets.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var secrets Secrets

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&secrets); err != nil {
		panic(err)
	}

	fmt.Println("Username:", secrets.ConnectionURI)

	var uri = secrets.ConnectionURI

	//connect to db
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

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
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	// configure apis
	router := gin.Default()

	//test APIs
	router.GET("/test", getTest)
	router.GET("/", endPoint)

	router.Run("localhost:8080")
}

func endPoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "message app backend"})
}

func getTest(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, testArr)
}
