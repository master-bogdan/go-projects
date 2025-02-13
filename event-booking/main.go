package main

import (
	"github.com/gin-gonic/gin"
	"github.com/master-bogdan/event-booking/db"
	"github.com/master-bogdan/event-booking/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
