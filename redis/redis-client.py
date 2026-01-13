import redis
import os
from dotenv import load_dotenv

# Load environment variables from .env file
load_dotenv()

# Get Redis credentials from environment variables
host = os.getenv('REDIS_HOST', '127.0.0.1')
port = int(os.getenv('REDIS_PORT', 6379))
password = os.getenv('REDIS_PASSWORD')
db = int(os.getenv('REDIS_DB', 0))

# Connect to Redis with authentication
r = redis.Redis(
    host=host,
    port=port,
    password=password,
    db=db,
    decode_responses=True
)

# Test connection
try:
    r.ping()
    print("✓ Connected to Redis successfully")
except redis.exceptions.AuthenticationError:
    print("✗ Authentication failed - invalid password")
except redis.exceptions.ConnectionError:
    print("✗ Could not connect to Redis server")

r.set('x', '42')
print(r.get('x'))  # '42'
