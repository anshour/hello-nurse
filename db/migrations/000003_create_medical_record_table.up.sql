CREATE TABLE IF NOT EXISTS medical_records (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    identity_number VARCHAR(16) NOT NULL UNIQUE,
    FOREIGN KEY (identity_number) REFERENCES patients(identity_number),
    symptoms VARCHAR(2000) NOT NULL, 
    medications VARCHAR(2000) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);