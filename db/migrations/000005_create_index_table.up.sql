CREATE INDEX idx_users_nip ON users (nip);
CREATE INDEX idx_users_id ON users (id);
CREATE INDEX idx_patients_id ON patients (id);
CREATE INDEX idx_medical_records_id ON medical_records (id);
CREATE INDEX idx_medical_records_patient_id ON medical_records (id);
