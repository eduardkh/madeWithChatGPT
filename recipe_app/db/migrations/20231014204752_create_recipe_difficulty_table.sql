-- +goose Up
-- +goose StatementBegin
CREATE TABLE RecipeDifficulty (
    recipe_id INTEGER,
    difficulty_id INTEGER,
    PRIMARY KEY (recipe_id, difficulty_id),
    FOREIGN KEY (recipe_id) REFERENCES Recipe(id),
    FOREIGN KEY (difficulty_id) REFERENCES Difficulty(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE RecipeDifficulty;
-- +goose StatementEnd
