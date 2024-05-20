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

# Create VLANs
vlan_commands = [
    'vlan 10',
    'name voice',
    'vlan 20',
    'name data',
    'vlan 30',
    'name servers',
]
output = conn.send_config_set(vlan_commands)
print(output)

# Verify VLAN configuration with TextFSM parsing
output = conn.send_command('show vlan brief', use_textfsm=True)
print(output)

# Close the connection
conn.disconnect()
