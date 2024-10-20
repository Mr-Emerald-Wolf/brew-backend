CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    uuid UUID DEFAULT gen_random_uuid(), 
    name VARCHAR(255) NOT NULL,   
    phone VARCHAR(20) NOT NULL UNIQUE,            
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE coffee (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE, 
    uuid UUID DEFAULT gen_random_uuid(), 
    name VARCHAR(255) NOT NULL,         
    origin VARCHAR(255) NOT NULL,      
    roast VARCHAR(255) NOT NULL,       
    process VARCHAR(255) NOT NULL,     
    price INTEGER NOT NULL,     
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, 
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ              
);
