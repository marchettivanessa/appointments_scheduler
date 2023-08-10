package handler

import (
	"appointments_scheduler/database"
	"appointments_scheduler/domain"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type BodyMessage struct {
	Message string `json:"message"`
}

// Get appointments handles the request and returns all appointments booked
func GetAppointments(c echo.Context, db *database.Database) error {
	appointmentDate := c.QueryParam("date")
	appointments, err := domain.GetConfirmedAppointments(appointmentDate, db)
	if err != nil {
		return fmt.Errorf("failed to get appointments from database: %w", err)
	}

	return c.JSON(http.StatusOK, appointments)
}

// CreateAppointment handles the request and returns , as a response, if the appointment was created correctly
func CreateAppointment(c echo.Context) error {
	var body BodyMessage
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed reading request body")
	}
	//aqui preciso salvar o body em um Appointment type
	return c.String(http.StatusOK, "Appointment Created")
}

// DeleteAppointment handles the request, and returns the response, that is an appointment deleted
func DeleteAppointment(c echo.Context, db *database.Database) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid ID")
	}
	err = domain.DeleteAppointmentById(idInt, db)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "failed deleting appointment from database")
	}

	return c.String(http.StatusOK, "Appointment deleted successfully")
}
