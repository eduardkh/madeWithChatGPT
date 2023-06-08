import ipaddress
import sys


def validate_input(ip_input):
    ip_parts = ip_input.split('/')
    if len(ip_parts) == 1:
        ip_parts = ip_input.split(' ')

    if len(ip_parts) != 2:
        return False

    ip_address = ip_parts[0].strip()
    subnet = ip_parts[1].strip()

    try:
        # Validate IP address
        ipaddress.IPv4Address(ip_address)
        # Validate subnet mask or CIDR
        ipaddress.IPv4Network(ip_address+'/'+subnet, strict=False)
        return True
    except (ipaddress.AddressValueError, ValueError):
        return False


def calculate_ip_details(ip_input):
    ip_parts = ip_input.split('/')
    if len(ip_parts) == 1:
        ip_parts = ip_input.split(' ')

    ip_address = ip_parts[0].strip()
    subnet = ip_parts[1].strip()

    ip_network = ipaddress.IPv4Network(ip_address+'/'+subnet, strict=False)

    address = ip_network.network_address
    broadcast = ip_network.broadcast_address
    host_range = f"{address + 1} - {broadcast - 1}"
    subnet_mask = ip_network.netmask
    host_number = ip_network.num_addresses - 2

    return {
        "Address": str(ip_network),
        "Network": str(address),
        "Broadcast": str(broadcast),
        "Host Range": host_range,
        "Subnet Mask": str(subnet_mask),
        "Host Number": host_number
    }


# Validate and process user input
if len(sys.argv) != 2:
    print("Please provide a valid IP address and subnet or CIDR.")
    sys.exit(1)

ip_input = sys.argv[1].strip()

if not validate_input(ip_input):
    print("Please provide a valid IP address and subnet or CIDR.")
    sys.exit(1)

# Calculate and display IP details
ip_details = calculate_ip_details(ip_input)

for key, value in ip_details.items():
    print(f"{key:<12}: {value}")
