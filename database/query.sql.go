// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package database

import (
	"context"
)

const createCoordinator = `-- name: CreateCoordinator :one
INSERT INTO coordinators(
email,firstName,lastName,password
) VALUES (?,?,?,?) RETURNING coordinatorid, firstname, lastname, email, password
`

type CreateCoordinatorParams struct {
	Email     string
	Firstname string
	Lastname  string
	Password  string
}

func (q *Queries) CreateCoordinator(ctx context.Context, arg CreateCoordinatorParams) (Coordinator, error) {
	row := q.db.QueryRowContext(ctx, createCoordinator,
		arg.Email,
		arg.Firstname,
		arg.Lastname,
		arg.Password,
	)
	var i Coordinator
	err := row.Scan(
		&i.Coordinatorid,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const createStudent = `-- name: CreateStudent :one
INSERT INTO students(
email,firstName,lastName,password
) VALUES (?,?,?,?) RETURNING studentid, email, firstname, lastname, password, supervisorid, projectid
`

type CreateStudentParams struct {
	Email     string
	Firstname string
	Lastname  string
	Password  string
}

func (q *Queries) CreateStudent(ctx context.Context, arg CreateStudentParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, createStudent,
		arg.Email,
		arg.Firstname,
		arg.Lastname,
		arg.Password,
	)
	var i Student
	err := row.Scan(
		&i.Studentid,
		&i.Email,
		&i.Firstname,
		&i.Lastname,
		&i.Password,
		&i.Supervisorid,
		&i.Projectid,
	)
	return i, err
}

const createSupervisor = `-- name: CreateSupervisor :one
INSERT INTO supervisors(
email,firstName,lastName,password
) VALUES (?,?,?,?) RETURNING supervisorid, firstname, lastname, email, password
`

type CreateSupervisorParams struct {
	Email     string
	Firstname string
	Lastname  string
	Password  string
}

func (q *Queries) CreateSupervisor(ctx context.Context, arg CreateSupervisorParams) (Supervisor, error) {
	row := q.db.QueryRowContext(ctx, createSupervisor,
		arg.Email,
		arg.Firstname,
		arg.Lastname,
		arg.Password,
	)
	var i Supervisor
	err := row.Scan(
		&i.Supervisorid,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const getAllStudents = `-- name: GetAllStudents :many
SELECT studentid, email, firstname, lastname, password, supervisorid, projectid FROM students
`

func (q *Queries) GetAllStudents(ctx context.Context) ([]Student, error) {
	rows, err := q.db.QueryContext(ctx, getAllStudents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Student
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.Studentid,
			&i.Email,
			&i.Firstname,
			&i.Lastname,
			&i.Password,
			&i.Supervisorid,
			&i.Projectid,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllSupervisors = `-- name: GetAllSupervisors :many
SELECT supervisorid, firstname, lastname, email, password FROM supervisors
`

func (q *Queries) GetAllSupervisors(ctx context.Context) ([]Supervisor, error) {
	rows, err := q.db.QueryContext(ctx, getAllSupervisors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Supervisor
	for rows.Next() {
		var i Supervisor
		if err := rows.Scan(
			&i.Supervisorid,
			&i.Firstname,
			&i.Lastname,
			&i.Email,
			&i.Password,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCoordinator = `-- name: GetCoordinator :one
SELECT coordinatorid, firstname, lastname, email, password FROM coordinators WHERE email =? LIMIT 1
`

func (q *Queries) GetCoordinator(ctx context.Context, email string) (Coordinator, error) {
	row := q.db.QueryRowContext(ctx, getCoordinator, email)
	var i Coordinator
	err := row.Scan(
		&i.Coordinatorid,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const getStudent = `-- name: GetStudent :one
SELECT studentid, email, firstname, lastname, password, supervisorid, projectid FROM students WHERE email =? LIMIT 1
`

func (q *Queries) GetStudent(ctx context.Context, email string) (Student, error) {
	row := q.db.QueryRowContext(ctx, getStudent, email)
	var i Student
	err := row.Scan(
		&i.Studentid,
		&i.Email,
		&i.Firstname,
		&i.Lastname,
		&i.Password,
		&i.Supervisorid,
		&i.Projectid,
	)
	return i, err
}

const getSupervisor = `-- name: GetSupervisor :one
SELECT supervisorid, firstname, lastname, email, password FROM supervisors WHERE email =? LIMIT 1
`

func (q *Queries) GetSupervisor(ctx context.Context, email string) (Supervisor, error) {
	row := q.db.QueryRowContext(ctx, getSupervisor, email)
	var i Supervisor
	err := row.Scan(
		&i.Supervisorid,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Password,
	)
	return i, err
}
