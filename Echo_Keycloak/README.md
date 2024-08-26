# Play with Keycloak

> Spin up a Keycloak server

```bash
docker run -d --name keycloak -p 8080:8080 \
  -e KEYCLOAK_ADMIN=admin \
  -e KEYCLOAK_ADMIN_PASSWORD=admin \
  quay.io/keycloak/keycloak:latest \
  start-dev
```

> Configure keycloak

- Create Realm (myrealm)
- Create Client (my-go-app)
- Create User

```bash
# Client ID - my-go-app
# Root URL - http://192.168.1.165:1323/
# Valid redirect URIs - http://192.168.1.165:1323/*
# Client authentication - On
```

```json
# Download adaptor configs
{
  "realm": "myrealm",
  "auth-server-url": "http://192.168.1.165:8080/",
  "ssl-required": "external",
  "resource": "my-go-app",
  "credentials": {
    "secret": "vdOwNrPhBMFfVUFIVFnGGwToGq8Xyfuq"
  },
  "confidential-port": 0
}
```

> Add AuthZ to Webapp

- Add roles under Client (my-go-app)
- Map roles to User
