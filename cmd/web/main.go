package main

import (
	"todo-list-golang/infrastructure"
	"todo-list-golang/internal/api"
)

func main() {
	// DB Connection
	db := infrastructure.NewDbConnection()
	defer db.Close()

	// HTTP
	routes := api.InitRoutes(db)
	routes.Run(":8080")
}