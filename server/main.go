package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	getEnv()

	switch os.Getenv("MODE") {
	case "DEBUG":
		gin.SetMode(gin.DebugMode)
	case "RELEASE":
		gin.SetMode(gin.ReleaseMode)
	case "TEST":
		gin.SetMode(gin.TestMode)
	}

	r := setRouter()

	r.Run(":" + os.Getenv("PORT"))
}
func getEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("A error occuered while opening .env file\n", err.Error())
	}
}
func setRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r
}
