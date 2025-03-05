-- name: CreateUser :one
INSERT INTO users(
userName,email,firstName,lastName,password
) VALUES (?,?,?,?,?) RETURNING *;
-- name: GetUser :one
SELECT * FROM users WHERE email =? LIMIT 1;
-- name: GetAllUsers :many
SELECT * FROM users;
-- name: GetAllTables :many
SELECT name FROM sqlite_schema;
-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;
-- name: UpdateUser :exec
UPDATE users 
SET username = ?, password = ?
WHERE id = ?;
