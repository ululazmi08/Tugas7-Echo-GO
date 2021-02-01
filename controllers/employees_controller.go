package controllers

import (
	"net/http"
	"pnp/echo-rest/models"

	"github.com/labstack/echo"
)

//FetchAllCustomers ...
func FetchAllEmployees(c echo.Context) (err error) {

	result, err := models.FetchEmployees()

	return c.JSON(http.StatusOK, result)
}

func AddEmployees(c echo.Context) (err error) {

	result, err := models.AddEmployee(c)

	return c.JSON(http.StatusOK, result)
}
