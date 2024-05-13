CREATE TABLE IF NOT EXISTS medical_records (
    identity_number VARCHAR(13) PRIMARY KEY NOT NULL,
    FOREIGN KEY (identity_number) REFERENCES patients(identity_number),
    symptoms VARCHAR(255) NOT NULL, 
    medications VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);