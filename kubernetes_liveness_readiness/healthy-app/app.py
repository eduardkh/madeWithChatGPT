from flask import Flask, jsonify
import os
import signal
import sys
import time

app = Flask(__name__)
start = time.time()


@app.route("/readyz")
def ready():
    # Pretend we need 3 s warm-up
    ready = time.time() - start > 3
    return ("OK", 200) if ready else ("Warming up", 503)


@app.route("/healthz")
def health():
    return jsonify(status="alive")


@app.route("/")
def root():
    return "I am healthy 🚑\n"

# Graceful shutdown for SIGTERM


def shutdown(sig, frame):
    print("SIGTERM → shutting down")
    sys.exit(0)


signal.signal(signal.SIGTERM, shutdown)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
