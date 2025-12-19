-- name: ListUsers :many
SELECT * FROM users
WHERE 
    (sqlc.narg('name_search')::TEXT IS NULL OR name ILIKE '%' || sqlc.narg('name_search')::TEXT || '%')
    AND (sqlc.narg('status')::user_status IS NULL OR status = sqlc.narg('status')::user_status)
ORDER BY
    CASE WHEN NOT sqlc.arg('order_desc')::BOOLEAN THEN created_at END DESC, -- Newest first.
    CASE WHEN sqlc.arg('order_desc')::BOOLEAN THEN created_at END ASC -- Oldest first.
LIMIT sqlc.arg('limit')::INT
OFFSET sqlc.arg('offset')::INT;

-- name: CountUsers :one
SELECT COUNT(*) FROM users
WHERE 
    (sqlc.narg('name_search')::TEXT IS NULL OR name ILIKE '%' || sqlc.narg('name_search')::TEXT || '%')
    AND status = 'ACTIVE';

-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: GetUserByEmailVerificationToken :one
SELECT * FROM users WHERE email_verification_token = $1 LIMIT 1;

-- name: GetUserByPasswordResetToken :one
SELECT * FROM users WHERE password_reset_token = $1 LIMIT 1;

-- name: CreateUser :exec
INSERT INTO users (
    name,
    email,
    password_hash
)
VALUES ($1, $2, $3);

-- name: UpdateUser :exec
UPDATE users
SET
    name = $2,
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateUserEmail :exec
UPDATE users
SET
    email = $2,
    email_verified = FALSE,
    updated_at = NOW()
WHERE id = $1;

-- name: VerifyUserEmail :exec
UPDATE users
SET
    email_verified = TRUE,
    email_verification_token = NULL,
    email_verification_expires = NULL,
    updated_at = NOW()
WHERE id = $1;

-- name: SetEmailVerificationToken :exec
UPDATE users
SET
    email_verification_token = $2,
    email_verification_expires = $3,
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE users
SET
    password_hash = $2,
    password_reset_token = NULL,
    password_reset_expires = NULL,
    last_password_change = NOW(),
    updated_at = NOW()
WHERE id = $1;

-- name: SetPasswordResetToken :exec
UPDATE users
SET
    password_reset_token = $2,
    password_reset_expires = $3,
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateLastLogin :exec
UPDATE users
SET 
    last_login = NOW(), 
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
