-- +goose Up
-- +goose StatementBegin
CREATE TABLE RecipeTechnique (
    recipe_id INTEGER,
    technique_id INTEGER,
    PRIMARY KEY (recipe_id, technique_id),
    FOREIGN KEY (recipe_id) REFERENCES Recipe(id),
    FOREIGN KEY (technique_id) REFERENCES Technique(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE RecipeTechnique;
-- +goose StatementEnd
