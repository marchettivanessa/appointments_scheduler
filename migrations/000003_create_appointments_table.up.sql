CREATE TABLE IF NOT EXISTS scheduler.appointments(
    id SERIAL PRIMARY KEY,
    patient_name VARCHAR NOT NULL,
    appointment_time TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    status BOOL,
    patient_id INT,
    FOREIGN KEY (patient_id) REFERENCES Patients(patient_id)
);