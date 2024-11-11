-- name: CreatePost :one
INSERT INTO posts (title, content, createdAt)
VALUES (?, ?, datetime('now'))
RETURNING *;

-- name: ListAllPosts :many
SELECT id, title, content, createdAt FROM posts ORDER BY id DESC;
