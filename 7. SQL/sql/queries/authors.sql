-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (
    name, bio
) VALUES ($1, $2)
RETURNING *;

-- name: UpdateAuthor :one
UPDATE authors
set name = $2,
    bio = $3
WHERE id = $1
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;

-- name: ClearAuthorsTable :exec
DELETE FROM authors;

-- name: ListAuthorBooks :many
-- SELECT authors.*, books.*
-- See https://docs.sqlc.dev/en/stable/howto/embedding.html
SELECT sqlc.embed(authors), sqlc.embed(books)
FROM authors
         JOIN books ON books.author_id = authors.id
WHERE authors.id = $1;