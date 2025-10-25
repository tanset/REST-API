package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tanset/REST-API/db"
	"github.com/tanset/REST-API/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}


