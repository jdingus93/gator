-- Schema definition
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT UNIQUE NOT NULL
);

-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
select * from users where name = $1;

-- name: DeleteAllUsers :exec
delete from users;

-- name: GetUsers :many
select * from users;

-- name: GetUserByName :one
SELECT * FROM users
WHERE name = $1
LIMIT 1;

-- name: GetUserByID :one
SELECT id, created_at, updated_at, name 
FROM users 
WHERE id = $1;