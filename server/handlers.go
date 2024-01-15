package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/paulojunior/trafilea/contract"
)

// @Summary Save number
// @Description Save a number to the collection
// @Accept json
// @Produce json
// @Param number body integer true "Number to be saved" format(number)
// @Success 201 {object} object "Number saved successfully"
// @Failure 400 {string} string "Failed to parse request body, number not found or not a valid string in the JSON body, or invalid number format"
// @Failure 500 {string} string "Internal server error"
// @Router /number [post]
func HandlerSaveNumber(c echo.Context, numberService contract.NumberService) error {
	var requestBody map[string]interface{}

	if err := c.Bind(&requestBody); err != nil {
		return c.String(http.StatusBadRequest, "Failed to parse request body")
	}

	param, ok := requestBody["number"].(string)
	if !ok || param == "" {
		return c.String(http.StatusBadRequest, "Number not found or not a valid string in the JSON body")
	}

	num, err := strconv.Atoi(param)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid number format")
	}

	err = numberService.Save(num)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error saving number"+err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "Number saved successfully"})
}

// @Summary Get number information
// @Description Retrieve information for a specific number
// @Produce json
// @Param number path int true "Number to retrieve information for" format(number)
// @Success 200 {object} object "Successfully retrieved number information"
// @Failure 400 {string} string "Invalid number format in the URL or number not found in the URL"
// @Failure 404 {string} string "Number not found"
// @Failure 500 {string} string "Internal server error"
// @Router /number/{number} [get]
func HandlerGetNumber(c echo.Context, numberService contract.NumberService) error {
	param := c.Param("number")

	if param == "" {
		return c.String(http.StatusBadRequest, "Number not found in the URL")
	}

	num, err := strconv.Atoi(param)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid number format in the URL")
	}

	number, err := numberService.FindNumber(num)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error fetching number: "+err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"number": number})
}

// @Summary Get all numbers
// @Description Retrieve information for all numbers
// @Produce json
// @Success 200 {array} string "Successfully retrieved all numbers"
// @Failure 500 {string} string "Internal server error"
// @Router /number [get]
func HandlerGetAllNumbers(c echo.Context, numberService contract.NumberService) error {
	numbers, err := numberService.FindAll()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, numbers)
}
