import os
import json
from flask import Flask, jsonify, render_template
from netmiko import ConnectHandler
from apscheduler.schedulers.background import BackgroundScheduler

app = Flask(__name__)
mac_addresses = []


def get_mac_addresses():
    global mac_addresses
    # Set the NET_TEXTFSM environment variable to point to the templates directory
    script_dir = os.path.dirname(os.path.abspath(__file__))
    templates_dir = os.path.join(
        script_dir, 'ntc-templates', 'ntc_templates', 'templates')
    os.environ['NET_TEXTFSM'] = templates_dir

    # SSH configuration
    device = {
        'device_type': 'cisco_ios',
        'host': '192.168.188.200',
        'username': 'cisco',
        'password': 'cisco',
    }

    # Create a connection
    conn = ConnectHandler(**device)

    # Run the command with TextFSM parsing
    output = conn.send_command(
        'show mac address-table vlan 20', use_textfsm=True)

    # Close the connection
    conn.disconnect()

    mac_addresses = output


@app.route('/')
def index():
    return render_template('index.html', mac_addresses=mac_addresses)


@app.route('/api/mac_addresses')
def api_mac_addresses():
    return jsonify(mac_addresses)


def scheduled_task():
    get_mac_addresses()


if __name__ == '__main__':
    scheduler = BackgroundScheduler()
    scheduler.add_job(scheduled_task, 'interval', seconds=5)
    scheduler.start()
    get_mac_addresses()
    app.run(debug=True)
