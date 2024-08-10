CREATE TYPE priority AS ENUM ('low', 'mid', 'high');

CREATE TABLE IF NOT EXISTS todos(
    id UUID PRIMARY KEY,
    body TEXT NOT NULL,
    priority priority NOT NULL DEFAULT('low'),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT (now()) NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT (now()) NOT NULL,
    deadline TIMESTAMP WITH TIME ZONE
);