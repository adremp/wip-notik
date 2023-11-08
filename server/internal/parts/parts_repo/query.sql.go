// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package parts_repo

import (
	"context"
)

const create = `-- name: Create :one
with part_el as (
	select
		coalesce(max(part_order), 0) + 1 as part_order
	from
		parts
	where
		page_id = $3
)
insert into
	parts (part_order, variant, title, page_id)
values
	(part_el.part_order, $1, $2, $3) RETURNING id, part_order, variant, title, page_id
`

type CreateParams struct {
	Variant PartType
	Title   string
	PageID  int64
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (Part, error) {
	row := q.db.QueryRow(ctx, create, arg.Variant, arg.Title, arg.PageID)
	var i Part
	err := row.Scan(
		&i.ID,
		&i.PartOrder,
		&i.Variant,
		&i.Title,
		&i.PageID,
	)
	return i, err
}