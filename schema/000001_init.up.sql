-- File: 000001_init.up.sql

-- Creating todo table
CREATE TABLE IF NOT EXISTS todo (
                                    id SERIAL PRIMARY KEY,
                                    title TEXT NOT NULL CHECK (LENGTH(title) <= 255),
                                    description TEXT,
                                    completed BOOLEAN NOT NULL DEFAULT FALSE,
                                    created_at TIMESTAMP
);