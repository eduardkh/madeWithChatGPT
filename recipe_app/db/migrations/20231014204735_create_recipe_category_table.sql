-- +goose Up
-- +goose StatementBegin
CREATE TABLE RecipeCategory (
    recipe_id INTEGER,
    category_id INTEGER,
    PRIMARY KEY (recipe_id, category_id),
    FOREIGN KEY (recipe_id) REFERENCES Recipe(id),
    FOREIGN KEY (category_id) REFERENCES Category(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE RecipeCategory;
-- +goose StatementEnd
