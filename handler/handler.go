package handler

import (
	"appointments_scheduler/database"
	"appointments_scheduler/domain"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// Get appointments handles the request and returns all appointments booked
func GetAppointments(c echo.Context, db *database.Database) error {
	// 1. Receber o request
	// 2. Handles the request sem criar régra de negócio.
	// 3. Chama a função com a regra de negócio
	appointmentDate := c.QueryParam("date")
	appointments, err := domain.GetConfirmedAppointments(appointmentDate, db)
	if err != nil {
		return fmt.Errorf("failed to get appointments from database: %w", err)
	}

	return c.JSON(http.StatusOK, appointments)
}

// CreateAppointment handles the request and returns , as a response, the appointment created
func CreateAppointment(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

// DeleteAppointment handles the request, and returns the response, that is an appointment deleted
func DeleteAppointment(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
