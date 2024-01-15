package server

import (
	"github.com/labstack/echo/v4"
	"github.com/paulojunior/trafilea/contract"
	_ "github.com/paulojunior/trafilea/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SetRoutes(e *echo.Echo, numberService contract.NumberService) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	g := e.Group("/api/v1")

	g.POST("/numbers", func(c echo.Context) error {
		return HandlerSaveNumber(c, numberService)
	})

	g.GET("/numbers/:number", func(c echo.Context) error {
		return HandlerGetNumber(c, numberService)
	})

	g.GET("/numbers", func(c echo.Context) error {
		return HandlerGetAllNumbers(c, numberService)
	})
}
