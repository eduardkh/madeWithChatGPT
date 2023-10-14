-- +goose Up
-- +goose StatementBegin
CREATE TABLE RecipeIngredient (
    recipe_id INTEGER,
    ingredient_id INTEGER,
    quantity VARCHAR(50) NOT NULL,
    PRIMARY KEY (recipe_id, ingredient_id),
    FOREIGN KEY (recipe_id) REFERENCES Recipe(id),
    FOREIGN KEY (ingredient_id) REFERENCES Ingredient(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE RecipeIngredient;
-- +goose StatementEnd
