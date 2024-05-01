package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type test struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var testArr = []test{
	{ID: "1", Title: "Blue Train"},
	{ID: "2", Title: "Jeru"},
}

func main() {
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
