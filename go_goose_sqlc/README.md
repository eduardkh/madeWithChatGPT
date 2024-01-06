# Utilize Goose and SQLc to work with postgres

> Spin up a test postgres container

```bash
docker run --rm -d \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -p 5432:5432 \
  postgres
```

> Install Goose and SQLc

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

> Create and apply a schema with Goose

```bash
# create the schema
goose -dir migrations create create_todos_table sql
# edit migrations/20240106141050_create_todos_table.sql file

# apply the schema
goose -dir migrations postgres "host=localhost user=postgres password=postgres dbname=postgres sslmode=disable" up
goose -dir migrations postgres "host=localhost user=postgres password=postgres dbname=postgres sslmode=disable" status
```

> SQLc

```bash
# populate the sqlc.yaml file
sqlc init

# create and populate the query/todos.sql file

# generate go code
sqlc generate

# run simple CRUD in main.go file
```
