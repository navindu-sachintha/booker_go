-- name: CreateBook :one
INSERT INTO books (id, isbn, name, author, publication, created_at, updated_at)
VALUES (
    $1,$2,$3,$4,$5,$6,$7
)
RETURNING *;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY name;

-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: UpdateBookName :exec
UPDATE books
SET name = $3, updated_at = $2
WHERE id = $1;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;