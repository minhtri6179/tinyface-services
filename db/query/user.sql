-- name: CreateUser :one
INSERT INTO user (
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