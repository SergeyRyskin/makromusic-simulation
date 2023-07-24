-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="Europe/Moscow";

-- Create user table
CREATE TABLE user (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    name VARCHAR (255) NOT NULL,
    surname VARCHAR (255) NOT NULL,
    username VARCHAR (255) NOT NULL,
    user_status INT NOT NULL,
    user_attrs JSONB NOT NULL
);

-- Add indexes
CREATE INDEX active_users ON users (title) WHERE user_status = 1;