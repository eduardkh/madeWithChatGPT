# recipe_app

> database

```bash
# go mod init recipe_app

# install goose package
go install github.com/pressly/goose/v3/cmd/goose@latest

# set environment variable
export GOOSE_DRIVER=sqlite3
export GOOSE_DBSTRING=./db/recipe.db

## create all tables
goose -dir db/migrations create create_category_table sql
goose -dir db/migrations create create_technique_table sql
goose -dir db/migrations create create_difficulty_table sql
goose -dir db/migrations create create_ingredient_table sql
goose -dir db/migrations create create_recipe_table sql
goose -dir db/migrations create create_recipe_category_table sql
goose -dir db/migrations create create_recipe_technique_table sql
goose -dir db/migrations create create_recipe_ingredient_table sql
goose -dir db/migrations create create_recipe_difficulty_table sql

goose -dir db/migrations status
# roll up migration
goose -dir db/migrations up
goose -dir db/migrations down
```

> populate the tables

```sql
-- Category
INSERT INTO Category (name) VALUES
    ('Appetizer'),
    ('Main Course'),
    ('Dessert'),
    ('Beverage'),
    ('Snack');

-- Technique
INSERT INTO Technique (name) VALUES
    ('Baking'),
    ('Frying'),
    ('Grilling'),
    ('Roasting'),
    ('Steaming');

-- Difficulty
INSERT INTO Difficulty (level) VALUES
    (1), (2), (3), (4), (5),
    (6), (7), (8), (9), (10);

-- Ingredient
INSERT INTO Ingredient (name, unit) VALUES
    ('Flour', 'cup'),
    ('Sugar', 'cup'),
    ('Salt', 'tsp'),
    ('Eggs', 'unit'),
    ('Milk', 'cup');

-- Recipe
INSERT INTO Recipe (slug, author, title, image, introduction, prep_time, cook_time, instructions, publish_date, create_date) VALUES
    ('chocolate-chip-cookies', 'Eddie', 'Chocolate Chip Cookies', 'image_url_here', 'Delicious homemade cookies.', 15, 12, 'Mix ingredients and bake at 350°F for 12 minutes.', '2023-10-14', '2023-10-14');

-- RecipeCategory
INSERT INTO RecipeCategory (recipe_id, category_id) VALUES
    (1, 3);

-- RecipeTechnique
INSERT INTO RecipeTechnique (recipe_id, technique_id) VALUES
    (1, 1);

-- RecipeIngredient
INSERT INTO RecipeIngredient (recipe_id, ingredient_id, quantity) VALUES
    (1, 1, '2'),
    (1, 2, '1'),
    (1, 3, '0.5'),
    (1, 4, '2'),
    (1, 5, '0.5');

-- RecipeDifficulty
INSERT INTO RecipeDifficulty (recipe_id, difficulty_id) VALUES
    (1, 4);
```

> query recipe

```sql
WITH
    RecipeData AS (
        SELECT * FROM Recipe WHERE slug = 'chocolate-chip-cookies'
    ),
    CategoryData AS (
        SELECT GROUP_CONCAT(Category.name) AS categories
        FROM Category
        JOIN RecipeCategory ON Category.id = RecipeCategory.category_id
        JOIN RecipeData ON RecipeData.id = RecipeCategory.recipe_id
    ),
    TechniqueData AS (
        SELECT GROUP_CONCAT(Technique.name) AS techniques
        FROM Technique
        JOIN RecipeTechnique ON Technique.id = RecipeTechnique.technique_id
        JOIN RecipeData ON RecipeData.id = RecipeTechnique.recipe_id
    ),
    IngredientData AS (
        SELECT GROUP_CONCAT(Ingredient.name || ' ' || RecipeIngredient.quantity || ' ' || Ingredient.unit) AS ingredients
        FROM Ingredient
        JOIN RecipeIngredient ON Ingredient.id = RecipeIngredient.ingredient_id
        JOIN RecipeData ON RecipeData.id = RecipeIngredient.recipe_id
    ),
    DifficultyData AS (
        SELECT Difficulty.level
        FROM Difficulty
        JOIN RecipeDifficulty ON Difficulty.id = RecipeDifficulty.difficulty_id
        JOIN RecipeData ON RecipeData.id = RecipeDifficulty.recipe_id
    )
SELECT
    RecipeData.*,
    CategoryData.categories,
    TechniqueData.techniques,
    IngredientData.ingredients,
    DifficultyData.level AS difficulty_level
FROM
    RecipeData
JOIN
    CategoryData ON 1=1
JOIN
    TechniqueData ON 1=1
JOIN
    IngredientData ON 1=1
JOIN
    DifficultyData ON 1=1;

```
