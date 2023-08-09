# fiber_casbin

> Retrieve all todos

```bash
# For an admin:
curl -H "X-User: adminUser" http://localhost:3000/todos

# For a regular user:
curl -H "X-User: regularUser" http://localhost:3000/todos
```

> Create a new todo

```bash
# For an admin:
curl -X POST -H "X-User: adminUser" -H "Content-Type: application/json" -d '{"id": 2, "title": "Another Todo", "is_done": false}' http://localhost:3000/todos

# For a regular user (this should fail based on our Casbin policy):
curl -X POST -H "X-User: regularUser" -H "Content-Type: application/json" -d '{"id": 2, "title": "Another Todo", "is_done": false}' http://localhost:3000/todos
```

> Update a todo

```bash
# For an admin:
curl -X PUT -H "X-User: adminUser" -H "Content-Type: application/json" -d '{"title": "Updated Todo", "is_done": true}' http://localhost:3000/todos/1

# For a regular user (this should fail based on our Casbin policy):
curl -X PUT -H "X-User: regularUser" -H "Content-Type: application/json" -d '{"title": "Updated Todo", "is_done": true}' http://localhost:3000/todos/1
```

> Delete a todo

```bash
# For an admin:
curl -X DELETE -H "X-User: adminUser" http://localhost:3000/todos/1

# For a regular user (this should fail based on our Casbin policy):
curl -X DELETE -H "X-User: regularUser" http://localhost:3000/todos/1
```
