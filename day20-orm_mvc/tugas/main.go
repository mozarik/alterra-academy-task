package main

import (
	"tugas-mvc/config"
	"tugas-mvc/routes"
)

func main() {
	config.InitDB()
	echoApp := routes.NewRoutes()
	err := echoApp.Start(":8080")
	if err != nil {
		panic(err)
	}
}
