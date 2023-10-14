-- +goose Up
-- +goose StatementBegin
CREATE TABLE Difficulty (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    level INTEGER CHECK (level BETWEEN 1 AND 10) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Difficulty;
-- +goose StatementEnd
