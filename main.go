package main

import (
	"etmarket/project/config"
	"etmarket/project/routes"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	//development
	e := echo.New()
	config.InitDb()
	config.InitPort()
	routes.New(e)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))

}
