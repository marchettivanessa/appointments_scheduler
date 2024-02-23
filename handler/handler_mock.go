package handler

import (
    "appointments_scheduler/database"
    "appointments_scheduler/domain"
)

type MockDBInterface struct{}

func (m *MockDBInterface) GetConfirmedAppointments(appointmentDate string, db *database.Database) ([]domain.Appointment, error) {
    appointments := []domain.Appointment{
        {ID: 1, PatientName: "João Gomes", AppointmentTime: "2024-02-19 10:00:00", CreatedAt: "2024-02-19 08:00:00", Status: "CONFIRMED", PatientID: 1},
        {ID: 2, PatientName: "Maria José", AppointmentTime: "2024-02-20 11:00:00", CreatedAt: "2024-02-19 09:00:00", Status: "CONFIRMED", PatientID: 2},
    }
    return appointments, nil
}

func (m *MockDBInterface) ValidateAppointment(appointment domain.Appointment, db *database.Database) (*domain.Appointment, error) {
    return &appointment, nil
}

func (m *MockDBInterface) DeleteAppointmentById(appointmentID int, db *database.Database) error {
    return nil
}
