CREATE TABLE IF NOT EXISTS shops (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    terms TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now() NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
);



-- CREATE TABLE IF NOT EXISTS categories (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
--     name VARCHAR(255) NOT NULL,
--     user_id UUID NOT NULL
-- );

