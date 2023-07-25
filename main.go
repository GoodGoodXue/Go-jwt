package main

import (
	"Practice/Go-Projects/jwt/initializes"
	"Practice/Go-Projects/jwt/models"
	"Practice/Go-Projects/jwt/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializes.ConnectToDb()
	models.SyncDatabase()
}

func main() {
	r := gin.Default()

	routes.ApiRoutes(r)

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	r.Run()

}
