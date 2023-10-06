-- name: CreateUser :one
INSERT INTO "user" (
        user_name,
        hashed_password,
        salt,
        first_name,
        last_name,
        date_of_birth,
        email
    )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
-- name: GetUserByUsername :one
SELECT *
FROM "user"
WHERE user_name = $1;
-- name: UpdateUser :one
UPDATE "user"
SET user_name = $1,
    hashed_password = $2,
    salt = $3,
    first_name = $4,
    last_name = $5,
    date_of_birth = $6,
    email = $7
WHERE id = $8
RETURNING *;
-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;