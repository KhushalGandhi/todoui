package main

import (
	"fmt"
	"todoui/models"
	"todoui/routers"
)

func main() {
	r := routers.RegisterRoutes()

	models.ConnecttoDatabase()

	fmt.Println("Successfully connected")
	r.Run("localhost:8080")
}
