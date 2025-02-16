// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package database

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(
userName,email,firstName,lastName,password
) VALUES (?,?,?,?,?) RETURNING id, username, email, firstname, lastname, password
`

type CreateUserParams struct {
	Username  string
	Email     string
	Firstname string
	Lastname  string
	Password  string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.Firstname,
		arg.Lastname,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Firstname,
		&i.Lastname,
		&i.Password,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, email, firstname, lastname, password FROM users WHERE email =? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Firstname,
		&i.Lastname,
		&i.Password,
	)
	return i, err
}
