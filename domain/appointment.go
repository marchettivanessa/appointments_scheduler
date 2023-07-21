package model

type Appointment struct {
	ID               int
	PatientName      Patient
	Time             string
	ContactOption    Patient
	CauseOfTreatment string
}

type Patient struct {
	ID          int
	Name        string
	DateOfBirth string
	Addres      string
	Email       string
	PhoneNumber string
	//ID do agendamento?
}


// função com a regra de negócio  do endpoint aqui. ex: query a database etc