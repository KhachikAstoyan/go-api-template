-- name: CreateUser :one
INSERT INTO users (username, email, password_hash, role)
VALUES ($1, $2, $3, $4)
RETURNING id, username, email, created_at;

-- name: GetUserByEmail :one
SELECT id, username, email, role, is_active 
FROM users 
WHERE email = $1;

-- name: UpdateUserLastLogin :exec
UPDATE users 
SET last_login = NOW() 
WHERE id = $1;

-- name: AuthenticateUser :one
SELECT id, password_hash 
FROM users 
WHERE email = $1 AND is_active = TRUE;

-- name: ListUsers :many
SELECT id, username, email, role, created_at
FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CountActiveUsers :one
SELECT COUNT(*) 
FROM users 
WHERE is_active = TRUE;