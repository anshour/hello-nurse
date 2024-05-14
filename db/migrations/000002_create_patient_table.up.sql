CREATE TABLE IF NOT EXISTS patients (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    identity_number VARCHAR(13) NOT NULL,
    phone_number VARCHAR(16) NOT NULL, 
    name VARCHAR(255) NOT NULL,
    birth_date TIMESTAMP NOT NULL,
    gender VARCHAR(6) NOT NULL,
    identity_image VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);