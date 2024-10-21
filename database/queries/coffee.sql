-- name: CreateCoffee :one
INSERT INTO coffee (user_id, uuid, name, origin, roast, process, price) 
VALUES ($1, gen_random_uuid(), $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetCoffeeByID :one
SELECT * FROM coffee 
WHERE id = $1;

-- name: GetCoffeeByUUID :one
SELECT * FROM coffee 
WHERE uuid = $1;

-- name: GetCoffeesByUserID :many
SELECT * FROM coffee 
WHERE user_id = $1;

-- name: UpdateCoffee :one
UPDATE coffee 
SET name = $1, origin = $2, roast = $3, process = $4, price = $5, updated_at = CURRENT_TIMESTAMP
WHERE uuid = $6
RETURNING *;

-- name: DeleteCoffee :exec
DELETE FROM coffee 
WHERE uuid = $1;

-- name: GetAllCoffees :many
SELECT * FROM coffee;

