package handler

import (
	"appointments_scheduler/domain"
	"net/http"

	"github.com/labstack/echo"
)

// Get appointments handles the request and returns all appointments booked
func GetAppointments(c echo.Context) error {
	// 1. Receber o request
	// 2. Handles the request sem criar régra de negócio.
	// 3. Chama a função com a regra de negócio
	appointments := domain.AllAppointments

	// id := c.Param("id")

	// db, err := database.NewDabataseWithMigrations(c database.Database)

	// // recebe a request. aqui eu lido com o request, mas sem criar regra de negócio aqui.
	// // o melhor é fazer em outra função, que posso chamar aqui
	// //retorna a expected response do endpoint, decidir o que retornar
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
