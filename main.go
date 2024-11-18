package main

import (
	"example.com/event-booking-api/db"
	"example.com/event-booking-api/routes"
	"github.com/gin-gonic/gin"
)

type MyInt int

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
