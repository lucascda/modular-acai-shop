-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: CreateUser :exec
INSERT INTO users (id, name, email, password)
VALUES ($1, $2, $3, $4);