-- name: CreateUser :one
INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET
    first_name = COALESCE($1, first_name),
    last_name = COALESCE($2, last_name),
    middle_name = COALESCE($3, middle_name),
    email = COALESCE($4, email),
    password_hash = COALESCE($5, password_hash),
    role = COALESCE($6, role),
    status = COALESCE($7, status),
    photo = COALESCE($8, photo),
    phone = COALESCE($9, phone),
    address = COALESCE($10, address),
    verified_token = COALESCE($11, verified_token),
    last_activity = COALESCE($12, last_activity),
    is_verified = COALESCE($13, is_verified),
    updated_at = COALESCE($14, updated_at)
WHERE id = $15
    RETURNING *;

-- name: GetOneByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;