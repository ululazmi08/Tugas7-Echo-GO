package controllers

import (
	"net/http"
	"pnp/echo-rest/models"

	"github.com/labstack/echo"
)

//FetchAllSuppliers ...
func FetchAllSuppliers(c echo.Context) (err error) {

	result, err := models.FetchSuppliers()

	return c.JSON(http.StatusOK, result)
}

//StoreSuppliers ...
func StoreSuppliers(c echo.Context) (err error) {

	result, err := models.StoreSuppliers(c)

	return c.JSON(http.StatusOK, result)
}

//UpdateSupplier ...
func UpdateSuppliers(c echo.Context) (err error) {

	result, err := models.UpdateSuppliers(c)

	return c.JSON(http.StatusOK, result)
}

//DeleteSuppliers ...
func DeleteSuppliers(c echo.Context) (err error) {

	result, err := models.DeleteSuppliers(c)

	return c.JSON(http.StatusOK, result)
}
