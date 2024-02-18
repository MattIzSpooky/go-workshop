-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateBook :one
INSERT INTO books (
    author_id, name, summary
) VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateBook :one
UPDATE books
set author_id = $2,
    name = $3,
    summary = $4
WHERE id = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;

-- name: ClearBooksTable :exec
DELETE FROM books;

-- name: GetBooksByAuthor :many
SELECT * FROM books
WHERE author_id = $1;