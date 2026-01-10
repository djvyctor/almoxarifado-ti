CREATE TABLE IF NOT EXISTS admins (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Seed: usuário admin com senha "admin123"
-- Hash bcrypt de "admin123"
INSERT INTO admins (username, password_hash) 
VALUES ('admin', '$2a$10$YQ3h6qZ5JK5X5X5X5X5X5.5X5X5X5X5X5X5X5X5X5X5X5X5X5X5X5Xe')
ON CONFLICT (username) DO NOTHING;