-- +goose Up
-- +goose StatementBegin
CREATE TABLE Recipe (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    slug VARCHAR(200) UNIQUE NOT NULL,
    author VARCHAR(100) NOT NULL,
    title VARCHAR(120) NOT NULL,
    image VARCHAR(500) NOT NULL,
    introduction TEXT,
    prep_time INTEGER NOT NULL,
    cook_time INTEGER NOT NULL,
    instructions TEXT NOT NULL,
    notes TEXT,
    publish_date DATETIME NOT NULL,
    create_date DATETIME NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Recipe;
-- +goose StatementEnd
