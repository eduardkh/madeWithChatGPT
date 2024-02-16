# Play with OPA

> OPA in a web application

```bash
# run webapp
go run .\main.go
# Access granted
curl --location 'http://localhost:8080/' --header 'User: admin'
# Access denied
curl --location 'http://localhost:8080/' --header 'User: notadmin'
```

> check permissions (OPA)

```bash
opa eval --data policy.rego --input input.json "data.myapp.authz.allow"
opa eval --data policy.rego --input input.json "data.myapp.authz.allow" --format raw
```
