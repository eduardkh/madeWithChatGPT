-- +goose Up
-- +goose StatementBegin
CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    task TEXT NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS todos;
-- +goose StatementEnd