package handler

import (
	"appointments_scheduler/database"
	"appointments_scheduler/domain"
	"appointments_scheduler/handler/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetAppointments(t *testing.T) {
	ctrl := gomock.NewController(t)
	dbMock := mocks.NewMockdbInterface(ctrl)
	handlerMock := NewHandler(dbMock)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/appointments?date=2024-02-19", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expectedResponse := domain.Appointment{
		ID:              1,
		PatientName:     "José da Silva",
		AppointmentTime: "2024-01-12 15:40:00",
		CreatedAt:       "2024-01-02 10:04:05.123456",
		Status:          "Confirmed",
		PatientID:       120,
	}
	dbMock.
		EXPECT().
		GetConfirmedAppointments(gomock.Any(), gomock.Any()).
		Return([]domain.Appointment{expectedResponse}, nil).
		Times(1)

	if assert.NoError(t, handlerMock.GetAppointments(c, &database.Database{})) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "José da Silva")
	}
}

func TestCreateAppointment(t *testing.T) {
	ctrl := gomock.NewController(t)
	dbMock := mocks.NewMockdbInterface(ctrl)
	handler := NewHandler(dbMock)

	e := echo.New()
	expectedJSON := `{"PatientName":"João Gomes","AppointmentTime":"2024-02-19 10:00:00","CreatedAt":"2024-02-19 08:00:00","Status":"CONFIRMED","PatientID":1}`
	req := httptest.NewRequest(http.MethodPost, "/appointments", strings.NewReader(expectedJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	var expectedScheduledAppointment domain.Appointment
	err := json.Unmarshal([]byte(expectedJSON), &expectedScheduledAppointment)
	if err != nil {
		t.Errorf("Failed to unmarshal expectedJSON: %v", err)
	}

	dbMock.
		EXPECT().
		ValidateAppointment(gomock.Any(), gomock.Any()).
		Return(&expectedScheduledAppointment, nil).
		Times(1)

	if assert.NoError(t, handler.CreateAppointment(c, &database.Database{})) {
		assert.Equal(t, http.StatusOK, rec.Code)
		expectedJSON, _ := json.Marshal(expectedScheduledAppointment)
		assert.Equal(t, string(expectedJSON), strings.TrimSpace(rec.Body.String()))
	}
}

func TestDeleteAppointment(t *testing.T) {
	ctrl := gomock.NewController(t)
	dbMock := mocks.NewMockdbInterface(ctrl)
	handler := NewHandler(dbMock)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/appointments/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	dbMock.
		EXPECT().
		DeleteAppointmentById(gomock.Any(), gomock.Any()).
		Return(nil).
		Times(1)

	if assert.NoError(t, handler.DeleteAppointment(c, &database.Database{})) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Appointment deleted successfully", rec.Body.String())
	}
}
