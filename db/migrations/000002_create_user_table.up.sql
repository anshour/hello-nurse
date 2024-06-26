CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    nip BIGINT NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL,
    password VARCHAR(255),
    role VARCHAR(5) NOT NULL,
    has_access BOOLEAN NOT NULL DEFAULT false,
    identity_image  VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);