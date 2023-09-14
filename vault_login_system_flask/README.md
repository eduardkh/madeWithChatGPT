# Vault login system Flask application

> Start Vault in Development Mode

```bash
docker run --cap-add=IPC_LOCK\
 -p 8200:8200\
 -e 'VAULT_DEV_LISTEN_ADDRESS=0.0.0.0:8200'\
 -e 'VAULT_DEV_ROOT_TOKEN_ID=root'\
 hashicorp/vault:1.14.2\
 vault server -dev
```

> Setup Vault for the app

```bash
# Export the Vault Address and Token
export VAULT_ADDR='http://127.0.0.1:8200'
export VAULT_TOKEN='root'  # This is the default root token for dev mode

# Enable UserPass Authentication Backend
vault auth enable userpass

# Upload the Policies to Vault
vault policy write alice-policy alice-policy.hcl
vault policy write bob-policy bob-policy.hcl

# Create Users (Alice and Bob) in Vault
vault write auth/userpass/users/alice password=alicepassword policies=alice-policy
vault write auth/userpass/users/bob password=bobpassword policies=bob-policy
```

> Setup Flask for the app

```bash
# Create a Virtual Environment and Activate It
python3 -m venv .venv
source .venv/bin/activate

# Install Necessary Libraries
pip install Flask hvac

# Run the Flask App
python app.py
```

> Test the Flask App with Curl

```bash
# Login as Alice
curl -X POST -H "Content-Type: application/json" -d '{"username":"alice", "password":"alicepassword"}' http://127.0.0.1:5000/login

# Access the Protected Route (using the token received from the login)
curl -H "Authorization: Bearer <TOKEN_FROM_LOGIN>" http://127.0.0.1:5000/protected/alice
```
