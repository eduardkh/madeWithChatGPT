from flask import Flask, request, jsonify
import hvac

app = Flask(__name__)

# Vault configuration
VAULT_ADDR = 'http://127.0.0.1:8200'
VAULT_TOKEN = 'root'  # This should ideally not be the root token in production
client = hvac.Client(url=VAULT_ADDR, token=VAULT_TOKEN)


@app.route('/login', methods=['POST'])
def login():
    data = request.get_json()
    username = data.get('username')
    password = data.get('password')

    try:
        # Authenticate using the userpass method
        login_response = client.auth.userpass.login(
            username=username, password=password)
        token = login_response['auth']['client_token']
        return jsonify({"message": "Authenticated", "token": token}), 200
    except Exception as e:
        return jsonify({"error": str(e)}), 400


@app.route('/protected/<username>', methods=['GET'])
def protected_resource(username):
    # Authorization header is in the format 'Bearer <token>'
    auth_header = request.headers.get('Authorization')
    if not auth_header:
        return jsonify({"error": "Token required"}), 401

    token = auth_header.split(" ")[1]
    client.token = token

    # Check if the token belongs to the user trying to access the resource
    token_data = client.lookup_token()
    if token_data['data']['meta']['username'] != username:
        return jsonify({"error": "Unauthorized"}), 403

    return jsonify({"message": f"This is {username}'s protected resource"}), 200


if __name__ == '__main__':
    app.run(host="0.0.0.0", debug=True)
