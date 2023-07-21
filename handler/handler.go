package handler

import (
	"appointments_scheduler/database"
	"net/http"

	"github.com/labstack/echo"
)

func GetAppointments(c echo.Context) error {

	id := c.Param("id")

	db, err := database.NewDabataseWithMigrations(c *database.Database)

	// recebe a request. aqui eu lido com o request, mas sem criar regra de negócio aqui.
	// o melhor é fazer em outra função, que posso chamar aqui
	//retorna a expected response do endpoint, decidir o que retornar
	return e.String(http.StatusOK, "OK")

}

func CreateAppointment(e echo.Context) error {
	return e.String(http.StatusOK, "OK")
}

func DeleteAppointment(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
