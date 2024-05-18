CREATE TABLE IF NOT EXISTS medical_records (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    identity_number BIGINT NOT NULL,
    FOREIGN KEY (identity_number) REFERENCES patients(identity_number),
    user_id uuid NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    symptoms VARCHAR(2000) NOT NULL, 
    medications VARCHAR(2000) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);