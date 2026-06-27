-- name: GetUserByEmail :one
SELECT * FROM client_users WHERE email = ? LIMIT 1;

-- name: CreateUser :exec
INSERT OR IGNORE INTO client_users (id, name, email, password_hash, project_status)
VALUES (?, ?, ?, ?, ?);
