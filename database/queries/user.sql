-- name: CreateUser :one
INSERT INTO users (name, phone, email, password_hash) 
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users 
WHERE id = $1;

-- name: GetUserByUUID :one
SELECT * FROM users 
WHERE uuid = $1;

-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1;

-- name: UpdateUser :one
UPDATE users 
SET name = $1, phone = $2, email = $3, updated_at = CURRENT_TIMESTAMP
WHERE uuid = $4
RETURNING *;

-- name: UpdateUserEmail :exec
UPDATE users 
SET email = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE uuid = $1;

-- name: GetAllUsers :many
SELECT * FROM users;