-- +goose Up
-- +goose StatementBegin
CREATE TABLE Technique (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(50) UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Technique;
-- +goose StatementEnd
