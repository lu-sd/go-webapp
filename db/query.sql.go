// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"
	"time"
)

const listAllPosts = `-- name: ListAllPosts :many
SELECT id, title, content, createdAt FROM posts ORDER BY id DESC
`

type ListAllPostsRow struct {
	ID        int64
	Title     string
	Content   string
	Createdat time.Time
}

func (q *Queries) ListAllPosts(ctx context.Context) ([]ListAllPostsRow, error) {
	rows, err := q.db.QueryContext(ctx, listAllPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAllPostsRow
	for rows.Next() {
		var i ListAllPostsRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.Createdat,
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
