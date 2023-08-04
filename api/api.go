package api

import (
	"appointments_scheduler/database"
	"appointments_scheduler/handler"

	"github.com/labstack/echo"
)

func RegisterHTTPRoutes(e *echo.Echo, db *database.Database) {
	h := func(c echo.Context) error {
		return handler.GetAppointments(c, db)
	}

	e.GET("/appointments", h)
	e.POST("/appointment", handler.CreateAppointment)
	// e.DELETE("/appointment/{id}", handler.DeleteAppointment)
}
