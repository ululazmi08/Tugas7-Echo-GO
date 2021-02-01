package routes

import (
	"pnp/echo-rest/controllers"

	"github.com/labstack/echo"
)

func SuppliersRoute(g *echo.Group) {

	g.POST("/list", controllers.FetchAllSuppliers)

	g.POST("/add", controllers.StoreSuppliers)

	g.POST("/update", controllers.UpdateSuppliers)

	g.POST("/delete", controllers.DeleteSuppliers)

}
