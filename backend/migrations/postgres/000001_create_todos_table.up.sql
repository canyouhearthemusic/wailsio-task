DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'priority') THEN
        CREATE TYPE priority AS ENUM ('low', 'mid', 'high');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS todos(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    body TEXT NOT NULL,
    priority priority NOT NULL DEFAULT('low'),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT (now()) NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT (now()) NOT NULL,
    deadline TIMESTAMP WITH TIME ZONE
);