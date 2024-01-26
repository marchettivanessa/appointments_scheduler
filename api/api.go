package api

import (
	"appointments_scheduler/database"
	"appointments_scheduler/domain"
	"appointments_scheduler/handler"

	"github.com/labstack/echo"
)

func RegisterHTTPRoutes(e *echo.Echo, db *database.Database) {
	nDomain := domain.Domain{}
	nHandler := handler.NewHandler(nDomain)
	getHandler := func(c echo.Context) error {
		return nHandler.GetAppointments(c, db)
	}

	postHandler := func(c echo.Context) error {
		return handler.CreateAppointment(c, db)
	}

	deleteHandler := func(c echo.Context) error {
		return handler.DeleteAppointment(c, db)
	}

	e.GET("/appointments", getHandler)
	e.POST("/appointments", postHandler)
	e.DELETE("/appointments/:id", deleteHandler)
}
