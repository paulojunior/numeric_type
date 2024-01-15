package main

import (
	"github.com/labstack/echo/v4"
	"github.com/paulojunior/trafilea/repository"
	"github.com/paulojunior/trafilea/server"
	"github.com/paulojunior/trafilea/services"
)

// @title Numeric Type API
// @version 1.0
// @description The NumericTypeService API categorizes and manages numeric values. The RESTful endpoints allow clients to store, retrieve, and query numeric values in memory.
// @contact.name Paulo Ferreira
// @contact.email jr@live.at
// @BasePath api/v1
func main() {

	// PART 1 - POC
	/*
		printers := services.NewPrinter()

		for i := 1; i <= 100; i++ {
			printer := printers.GetPrinter(i)
			fmt.Println(printer.PrintNumber(i))
		}
	*/

	// PART 2 - REST API
	e := echo.New()

	numberRepository := repository.NewNumberRepository()
	printers := services.NewPrinter()
	numberService := services.NewNumberService(numberRepository, printers)

	server.SetRoutes(e, numberService)

	e.Logger.Fatal(e.Start(":8080"))

}
