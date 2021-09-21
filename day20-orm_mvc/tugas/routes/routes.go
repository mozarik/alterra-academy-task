package routes

import (
	"tugas-mvc/controllers"

	"github.com/labstack/echo/v4"
)

func NewRoutes() *echo.Echo {
	echoInit := echo.New()

	echoInit.GET("/user", controllers.GetAllUsers)
	echoInit.POST("/user", controllers.AddUser)
	echoInit.GET("/user/:id", controllers.GetUser)

	return echoInit
}
