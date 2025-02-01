-- +goose Up
CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    middle_name VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(255) DEFAULT 'USER' NOT NULL,
    status VARCHAR(255) DEFAULT 'ACTIVE' NOT NULL,
    photo VARCHAR(255),
    phone VARCHAR(255),
    address VARCHAR(255),
    verified_token VARCHAR(255),
    last_activity VARCHAR(255) NOT NULL DEFAULT now(),
    is_verified BOOLEAN DEFAULT false NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE users;

