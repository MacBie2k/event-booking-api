package main

import (
	"github.com/MacBie2k/event-booking-api/db"
	"github.com/MacBie2k/event-booking-api/routes"
	"github.com/gin-gonic/gin"
)

type MyInt int

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
