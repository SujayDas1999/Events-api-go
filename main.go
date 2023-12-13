package main

import (
	database "events-api/helpers/db"
	"events-api/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	//var db database.DB

	_ = database.InitDB()

	//db.SetDB(postgresInstance)

	server := gin.Default()

	routes.EventRoutes(server)

	err := server.Run(":8080")

	if err != nil {
		fmt.Println(err)
	}

}
