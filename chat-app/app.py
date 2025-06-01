from flask import Flask, render_template
from flask_socketio import SocketIO
import redis

app = Flask(__name__)
app.config['SECRET_KEY'] = 'your_secret_key'
socketio = SocketIO(app)

# Connect to Redis
r = redis.Redis(host='localhost', port=6379)


@app.route('/')
def index():
    return render_template('index.html')


@socketio.on('message')
def handle_message(msg):
    # Publish message to Redis channel
    r.publish('chat', msg)
    # Emit the message to all clients
    socketio.emit('message', msg)


if __name__ == '__main__':
    socketio.run(app, host='0.0.0.0')
