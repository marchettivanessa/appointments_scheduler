package api

import (
	"appointments_scheduler/database"
	"appointments_scheduler/handler"

	"github.com/labstack/echo"
)

func RegisterHTTPRoutes(e *echo.Echo, db *database.Database) {
	e.GET("/appointments", handler.GetAppointments)
	e.POST("/appointment", handler.CreateAppointment)
	// e.DELETE("/appointment/{id}", handler.DeleteAppointment)
}

// TODO descobrir como chamar a db, sendo que o e.Get não suporta parametros, ja que mexeria com a assinatura dessa funcao no echo.
