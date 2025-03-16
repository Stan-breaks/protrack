-- name: CreateStudent :one
INSERT INTO students(
email,firstName,lastName,password
) VALUES (?,?,?,?) RETURNING *;
-- name: GetStudent :one
SELECT * FROM students WHERE email =? LIMIT 1;
-- name: GetStudentById :one
SELECT * FROM students WHERE studentId =? LIMIT 1;
-- name: GetAllStudents :many
SELECT * FROM students;
-- name: CreateCoordinator :one
INSERT INTO coordinators(
email,firstName,lastName,password
) VALUES (?,?,?,?) RETURNING *;
-- name: GetCoordinator :one
SELECT * FROM coordinators WHERE email =? LIMIT 1;
-- name: GetCoordinatorById :one
SELECT * FROM coordinators WHERE coordinatorId =? LIMIT 1;
-- name: CreateSupervisor :one
INSERT INTO supervisors(
email,firstName,lastName,password
) VALUES (?,?,?,?) RETURNING *;
-- name: GetSupervisor :one
SELECT * FROM supervisors WHERE email =? LIMIT 1;
-- name: GetSupervisorById :one
SELECT * FROM supervisors WHERE supervisorID =? LIMIT 1;
-- name: GetAllSupervisors :many
SELECT * FROM supervisors;
-- name: CreateProject :one
INSERT INTO projects(name,description
) VALUES (?,?)RETURNING *;
-- name: GetProject :one
SELECT * FROM projects WHERE projectId =? LIMIT 1;
-- name: AssignSupervisor :exec
UPDATE students SET supervisorId =? WHERE studentId =?;
