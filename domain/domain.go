package domain

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
	//ID do agendamento?
}

// função com a regra de negócio  do endpoint aqui. ex: query a database etc
func AllAppointments() []Appointment {
	// TODO: chamar a função, mas entender pq o erro com repository
	allAppointments := Repository.GetConfirmedAppointments
	// tratar as apoitments recebidas
	
	return allAppointments()

}
