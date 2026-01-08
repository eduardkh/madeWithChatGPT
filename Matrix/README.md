# Matrix Synapse + PostgreSQL (Podman / Docker Compose)

Production-ready Matrix Synapse deployment backed by PostgreSQL, running with **Podman Compose**.

This README reflects a **working, verified setup**.

---

## Prerequisites

- Podman
- podman-compose
- Port **8008** available
- Local DNS or `/etc/hosts` entry for your server name (example: `rocky9.local`)

---

## Directory Layout

```
.
├── docker-compose.yml
├── synapse-data/
│   ├── homeserver.yaml
│   ├── log.config
│   ├── rocky9.local.signing.key
│   └── media/
└── postgres-data/
```

---

## docker-compose.yml

```yaml
version: "3.7"

services:
  db:
    image: postgres:17
    container_name: matrix_db
    restart: unless-stopped
    environment:
      POSTGRES_DB: synapse
      POSTGRES_USER: synapse
      POSTGRES_PASSWORD: yourpassword
    volumes:
      - ./postgres-data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U synapse -d synapse"]
      interval: 5s
      timeout: 5s
      retries: 10

  synapse:
    image: matrixdotorg/synapse:latest
    container_name: matrix_synapse
    restart: unless-stopped
    depends_on:
      db:
        condition: service_healthy
    ports:
      - "8008:8008"
    volumes:
      - ./synapse-data:/data
```

---

## Step 1 – Create Directories

```bash
mkdir -p synapse-data postgres-data
```

---

## Step 2 – Generate Synapse Config & Signing Keys

```bash
podman run --rm   -v $(pwd)/synapse-data:/data   -e SYNAPSE_SERVER_NAME=rocky9.local   -e SYNAPSE_REPORT_STATS=no   matrixdotorg/synapse:latest generate
```

---

## Step 3 – Create Logging Configuration

```bash
cat <<'EOF' > synapse-data/log.config
version: 1

formatters:
  precise:
    format: "%(asctime)s - %(name)s - %(lineno)d - %(levelname)s - %(request)s - %(message)s"

handlers:
  file:
    class: logging.handlers.RotatingFileHandler
    formatter: precise
    filename: /data/homeserver.log
    maxBytes: 104857600
    backupCount: 10
    encoding: utf8

loggers:
  synapse:
    level: INFO
    handlers: [file]
    propagate: false

root:
  level: INFO
  handlers: [file]
EOF
```

---

## Step 4 – Minimal Working homeserver.yaml

```yaml
server_name: "rocky9.local"
public_baseurl: "http://rocky9.local:8008/"
report_stats: false
pid_file: /homeserver.pid
log_config: "/data/log.config"

media_store_path: "/data/media"

listeners:
  - port: 8008
    bind_addresses: ["0.0.0.0"]
    type: http
    tls: false
    x_forwarded: false
    resources:
      - names: [client]
        compress: true
      - names: [federation]
        compress: false

database:
  name: psycopg2
  args:
    user: synapse
    password: yourpassword
    database: synapse
    host: matrix_db
    port: 5432

enable_registration: true
enable_registration_without_verification: true

macaroon_secret_key: "REPLACE_ME"
signing_key_path: "/data/rocky9.local.signing.key"

trusted_key_servers:
  - server_name: matrix.org

password_config:
  enabled: true
```

---

## Step 5 – Fix Permissions

```bash
sudo chown -R 991:991 synapse-data
sudo chmod 755 synapse-data
sudo chmod 600 synapse-data/rocky9.local.signing.key
```

---

## Step 6 – Start Services

```bash
sudo podman compose up -d
```

---

## Step 7 – Verify

```bash
curl http://localhost:8008/_matrix/client/versions
```

---

## Step 8 – Create Admin User

```bash
sudo podman exec -it matrix_synapse   register_new_matrix_user   -c /data/homeserver.yaml   -u admin   -p 'StrongPasswordHere'   --admin
```

---

## Status

✅ Synapse running  
✅ PostgreSQL healthy  
✅ Logging works  
✅ Admin user created
