from flask import Flask
app = Flask(__name__)


@app.route("/")
def root():
    return "No probes here 🤷‍♂️\n"


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
