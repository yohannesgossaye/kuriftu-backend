-- name: CreateUser :one
INSERT INTO users (
    first_name, last_name, email, password_hash, phone, user_type
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id, first_name, last_name, email, phone, user_type, created_at, updated_at, last_login_at, is_active;