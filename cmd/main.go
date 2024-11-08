package main

import (
	"github.com/labstack/echo/v4"
	"localhost/api"
	userservice "localhost/services"
	"localhost/infrastructure"
)

func main() {
	e := echo.New()

	e.GET("/health-check", func(c echo.Context) error {
		return c.String(200, "OK")
	})


	userRepository := intrastructure.NewUserRepository()
	userService := userservice.NewUserService(userRepository)

	_ = userService

	userAPI := e.Group("/user")

	userController := api.NewUserAPI(userService)

	userAPI.GET("", userController.GetUser)
	userAPI.POST("", userController.CreateUser)
	userAPI.PUT("", userController.UpdateUser)
	userAPI.DELETE("", userController.DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}
