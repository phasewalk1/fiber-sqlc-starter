-- name: CreateUser :one
INSERT INTO users (
    fname, lname, email, salt, hash
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET
    fname = $1,
    lname = $2,
    email = $3,
    updated_at = NOW()
WHERE
    id = $4
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
