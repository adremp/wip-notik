// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package users_repo

import (
	"context"
)

const create = `-- name: Create :one
insert into
	users (username, email, password)
values
	($1, $2, $3) RETURNING id, username, email
`

type CreateParams struct {
	Username string
	Email    string
	Password string
}

type CreateRow struct {
	ID       int32
	Username string
	Email    string
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (CreateRow, error) {
	row := q.db.QueryRow(ctx, create, arg.Username, arg.Email, arg.Password)
	var i CreateRow
	err := row.Scan(&i.ID, &i.Username, &i.Email)
	return i, err
}

const delete = `-- name: Delete :exec
delete from
	users
where
	id = $1
`

func (q *Queries) Delete(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, delete, id)
	return err
}

const getByFields = `-- name: GetByFields :many
select id, username, email, password, created_at from users 
where (id = COALESCE(NULLIF($1::int, 0), id)) AND 
(username = COALESCE(NULLIF($2::text, ''), username)) AND 
(email = COALESCE(NULLIF($3::text, ''), email)) 
limit COALESCE(NULLIF($4::int, 0), 1)
`

type GetByFieldsParams struct {
	ID       int32
	Username string
	Email    string
	Limits   int32
}

func (q *Queries) GetByFields(ctx context.Context, arg GetByFieldsParams) ([]User, error) {
	rows, err := q.db.Query(ctx, getByFields,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.Limits,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
