package domain

import (
	"appointments_scheduler/database"
	"errors"
	"fmt"
)

type Appointment struct {
	ID              int
	PatientName     string
	AppointmentTime string
	CreatedAt       string
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


type Domain struct {}

func (d Domain) GetConfirmedAppointments(appointmentDate string, db *database.Database) ([]Appointment, error) {
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

func (d Domain) DeleteAppointmentById(appointmetnID int, db *database.Database) error {
	query := `DELETE * FROM appointments WHERE id = $1`

	_, err := db.Connection.Exec(query, appointmetnID)
	if err != nil {
		return fmt.Errorf("failed to delete the appointment %w", err)
	}
	return nil
}

func (d Domain) ValidateAppointment(appointment Appointment, db *database.Database) (*Appointment, error) {
	patientName := appointment.PatientName
	if patientName == "" {
		return nil, errors.New("not possible to save an appointment without the patient name")
	}

	appointmentTime := appointment.AppointmentTime
	if appointmentTime == "" {
		return nil, errors.New("not possible to save an appointment without the appointment time")
	}

	createdAt := appointment.CreatedAt
	if createdAt == "" {
		return nil, errors.New("not possible to save an appointment without the created at")
	}

	status := appointment.Status
	patientID := appointment.PatientID

	query := `
	INSERT INTO scheduler.appointments (patient_name, appointment_time, created_at, status, patient_id) VALUES ($1, $2< $3, $4, $5)
	`

	row := db.Connection.QueryRow(query, patientName, appointmentTime, createdAt, status, patientID)

	var insertedAppointment Appointment
	if err := row.Scan(&insertedAppointment.PatientName, &insertedAppointment.AppointmentTime, &insertedAppointment.CreatedAt, &insertedAppointment.Status, &insertedAppointment.PatientID); err != nil {
		return nil, errors.New("error inserting appointment into database")
	}

	return &insertedAppointment, nil
}
