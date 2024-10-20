-- name: CreateUser :one
INSERT INTO users (uuid, email, password_hash) 
VALUES (gen_random_uuid(), $1, $2)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users 
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1;

-- name: UpdateUserEmail :exec
UPDATE users 
SET email = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1;
