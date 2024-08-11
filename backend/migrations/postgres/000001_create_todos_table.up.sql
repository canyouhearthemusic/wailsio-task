DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'priority') THEN
        CREATE TYPE priority AS ENUM ('low', 'mid', 'high');
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status') THEN
        CREATE TYPE status AS ENUM ('in-progress', 'completed');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS todos(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    body TEXT NOT NULL,
    priority priority NOT NULL DEFAULT('low'),
    status status NOT NULL DEFAULT('in-progress'),
    deadline TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT (now()) NOT NULL,
    updated_at TIMESTAMP DEFAULT (now()) NOT NULL
);
