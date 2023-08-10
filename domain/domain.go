package domain

import (
	"appointments_scheduler/database"
	"fmt"
)

type Appointment struct {
	ID              int
	PatientName     Patient
	AppointmentTime string
	CreatedAt       string
	ContactOption   Patient
	Status          string
	PatientID       int
}

type Patient struct {
	ID               int
	Name             string
	DateOfBirth      string
	Addres           string
	Email            string
	CauseOfTreatment string
	PhoneNumber      string
}

func GetConfirmedAppointments(appointmentDate string, db *database.Database) ([]Appointment, error) {
	queryResults := []Appointment{}

	query := `
SELECT * FROM appointments WHERE status = 'CONFIRMED' AND appointment_time <= $1`

	err := db.Connection.Select(&queryResults, query, appointmentDate)
	if err != nil {
		return nil, fmt.Errorf("failed to select all appointments in this time range: %w", err)
	}

	if len(queryResults) == 0 {
		return nil, nil
	}

	return queryResults, nil
}

func DeleteAppointmentById(appointmetnID int, db *database.Database) error {
	query := `DELETE * FROM appointments WHERE id = $1`

	_, err := db.Connection.Exec(query, appointmetnID)
	if err != nil {
		return fmt.Errorf("failed to delete the appointment %w", err)
	}
	return nil
}
