-- name: CreateUser :one
INSERT INTO users(
userName,email,firstName,lastName,password
) VALUES (?,?,?,?,?) RETURNING *;
-- name: GetUser :one
SELECT * FROM users WHERE email =? LIMIT 1;
