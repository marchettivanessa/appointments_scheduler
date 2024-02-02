package handler

import (
	"appointments_scheduler/database"
	"appointments_scheduler/domain"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type BodyMessage struct {
	Message string `json:"message"`
}

type Handler struct {
	dbMethods dbInterface
}

func NewHandler(dbMethods dbInterface) Handler{
	return Handler{
		dbMethods: dbMethods,
	}
}

type dbInterface interface{
	GetConfirmedAppointments(string, *database.Database)([]domain.Appointment, error)
}

// Get appointments handles the request and returns all appointments booked
func (h Handler) GetAppointments(c echo.Context, db *database.Database) error {
	appointmentDate := c.QueryParam("date")
	appointments, err := h.dbMethods.GetConfirmedAppointments(appointmentDate, db)
	if err != nil {
		return fmt.Errorf("failed to get appointments from database: %w", err)
	}

	return c.JSON(http.StatusOK, appointments)
}

// CreateAppointment handles the request and returns , as a response, if the appointment was created correctly
func CreateAppointment(c echo.Context, db *database.Database) error {
	var body domain.Appointment
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed reading request body")
	}

	_, err := domain.ValidateAppointment(body, db)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, body)
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
