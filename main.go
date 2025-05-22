package main

import (
	"inventory-system/config"
	"inventory-system/routes"
)

func main() {
	config.InitDB()
	r := routes.SetupRoutes()
	r.Run(":8080")
}