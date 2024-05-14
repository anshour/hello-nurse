CREATE TABLE IF NOT EXISTS medical_records (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id uuid NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES patients(id),
    symptoms VARCHAR(255) NOT NULL, 
    medications VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);