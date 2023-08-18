package api

import (
	"appointments_scheduler/database"
	"appointments_scheduler/handler"

	"github.com/labstack/echo"
)

func RegisterHTTPRoutes(e *echo.Echo, db *database.Database) {
	getHandler := func(c echo.Context) error {
		return handler.GetAppointments(c, db)
	}

	postHandler := func(c echo.Context) error {
		return handler.CreateAppointment(c, db)
	}

	deleteHandler := func(c echo.Context) error {
		return handler.DeleteAppointment(c, db)
	}

	e.GET("/appointments", getHandler)
	e.POST("/appointment", postHandler)
	e.DELETE("/appointment/:id", deleteHandler)
}
