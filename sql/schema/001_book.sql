-- +goose Up

CREATE TABLE books (
    id UUID PRIMARY KEY,
    isbn VARCHAR(17) NOT NULL,
    name TEXT NOT NULL,
    author TEXT NOT NULL,
    publication TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE books;