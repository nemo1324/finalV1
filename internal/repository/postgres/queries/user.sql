-- name: CreateUser :one
INSERT INTO users (name, login, pass, status)
VALUES (@name, @login, @pass, @status)
RETURNING id;

-- name: GetUserByLogin :one
SELECT * FROM users
WHERE login = @login
LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = @id
LIMIT 1;
