package main

import (
	"gin-api/database"
	"gin-api/routes"
)

func main()  {
	database.StartDb()
	routes.HandleRequest()
}