package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"golang-rest-api-validation-example/controller"
	_ "golang-rest-api-validation-example/docs"
	"golang-rest-api-validation-example/handler"
	"golang-rest-api-validation-example/util"
)

// @title Golang User REST API
// @description Provides access to the core features of Golang User REST API
// @version 1.0
// @termsOfService http://swagger.io/terms/
// license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
func main() {

	e := echo.New()
	e.HTTPErrorHandler = handler.ErrorHandler
	userController := controller.NewUserController()
	e.Validator = util.NewValidationUtil()
	v1 := e.Group("/api/v1")
	{
		v1.GET("/users", userController.GetAllUser)
		v1.POST("/users", userController.SaveUser)
		v1.GET("/users/:id", userController.GetUser)
		v1.PUT("/users/:id", userController.UpdateUser)
		v1.DELETE("/users/:id", userController.DeleteUser)

	}

	e.GET("/api", controller.RedirectIndexPage)
	e.GET("/api/*", echoSwagger.WrapHandler)
	// echo server 9000 de başlatıldı.
	e.Logger.Fatal(e.Start(":9000"))
}
