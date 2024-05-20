import os
from netmiko import ConnectHandler

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
output = conn.send_command('show mac address-table', use_textfsm=True)

# Print the parsed output
# print(output[0]["uptime"])
print(output)

# Close the connection
conn.disconnect()
