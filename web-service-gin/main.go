package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"

	"encoding/json"
	"os"
)

type test struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var testArr = []test{
	{ID: "1", Title: "Pulp Fiction"},
	{ID: "2", Title: "No Country for Old Men"},
}

type Secrets struct {
	ConnectionURI string `json:"connectionURI"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	// configure apis
	router := gin.Default()

	//test APIs
	router.GET("/test", getTest)
	router.GET("/", endPoint)
	router.GET("/ws", webSocket)

	router.Run("localhost:8080")
}

func endPoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "message app backend"})
}

func getTest(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, testArr)
}

func webSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	for {
		conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
		time.Sleep(time.Second)
	}
}
