CREATE TABLE IF NOT EXISTS admins (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Força o email a ser salvo em minusculo, e valida
    email TEXT NOT NULL UNIQUE CHECK (
        email = lower(email) AND 
        email ~* '^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$'
    ),
    
    password_hash TEXT NOT NULL,
    
    -- Para lidar corretamente com fusos horários
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    -- Quando foi alterado por ultimo
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);