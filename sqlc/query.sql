-- name: CreateUser :one
INSERT INTO students(
email,firstName,lastName,password
) VALUES (?,?,?,?) RETURNING *;
-- name: GetStudent :one
SELECT * FROM students WHERE email =? LIMIT 1;
-- name: GetAllStudents :many
SELECT * FROM students;
-- name: CreateCoordinator :one
INSERT INTO coordinators(
email,firstName,lastName,password
) VALUES (?,?,?,?) RETURNING *;
-- name: GetCoordinator :one
SELECT * FROM coordinators WHERE email =? LIMIT 1;
-- name: CreateSupervisor :one
INSERT INTO supervisors(
email,firstName,lastName,password
) VALUES (?,?,?,?) RETURNING *;
-- name: GetSupervisor :one
SELECT * FROM supervisors WHERE email =? LIMIT 1;
-- name: GetAllSupervisors :many
SELECT * FROM supervisors;
