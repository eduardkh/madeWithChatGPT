# recipe_app

> database

```bash
# go mod init recipe_app

# install goose package
go install github.com/pressly/goose/v3/cmd/goose@latest

# set environment variable
export GOOSE_DRIVER=sqlite3
export GOOSE_DBSTRING=./db/recipe.db

# create all tables
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
