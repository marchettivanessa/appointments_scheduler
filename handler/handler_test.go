package handler

import (
	"appointments_scheduler/database"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetAppointments(t *testing.T) {
	dbMock := &MockDBInterface{}
	handler := NewHandler(dbMock)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/appointments?date=2024-02-19", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.GetAppointments(c, &database.Database{})) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "João Gomes")
		assert.Contains(t, rec.Body.String(), "Maria José")
	}
}

func TestCreateAppointment(t *testing.T) {
	dbMock := &MockDBInterface{}
	handler := NewHandler(dbMock)

	e := echo.New()
	appointmentJSON := `{"PatientName":"João Gomes","AppointmentTime":"2024-02-19 10:00:00","CreatedAt":"2024-02-19 08:00:00","Status":"CONFIRMED","PatientID":1}`
	req := httptest.NewRequest(http.MethodPost, "/appointments", strings.NewReader(appointmentJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.CreateAppointment(c, &database.Database{})) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "João Gomes")
	}
}

func TestDeleteAppointment(t *testing.T) {
	dbMock := &MockDBInterface{}
	handler := NewHandler(dbMock)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/appointments/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, handler.DeleteAppointment(c, &database.Database{})) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Appointment deleted successfully", rec.Body.String())
	}
}
