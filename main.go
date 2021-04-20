package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	_ "golang-rest-api-validation-example/docs"
	"golang-rest-api-validation-example/handler"
	"golang-rest-api-validation-example/routes"
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
	e.Validator = util.NewValidationUtil()

	routes.GetUserApiRoutes(e)
	routes.GetSwaggerRoutes(e)
	// echo server 9000 de başlatıldı.
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", util.GetEnv("SERVER_PORT", "9000"))))
}
