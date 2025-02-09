import socket
import struct


def parse_arp(packet):
    # Skip Ethernet header (14 bytes)
    arp_header = packet[14:42]
    arp = struct.unpack("!HHBBH6s4s6s4s", arp_header)
    return {
        'hw_type': arp[0],
        'proto_type': arp[1],
        'hw_size': arp[2],
        'proto_size': arp[3],
        'opcode': arp[4],
        'src_mac': arp[5],
        'src_ip': socket.inet_ntoa(arp[6]),
        'dst_mac': arp[7],
        'dst_ip': socket.inet_ntoa(arp[8])
    }


def format_mac(mac_bytes):
    return ':'.join(f"{b:02x}" for b in mac_bytes)


def main():
    s = socket.socket(socket.AF_PACKET, socket.SOCK_RAW, socket.ntohs(0x0003))
    s.bind(("eth0", 0))  # Change interface if needed

    while True:
        packet = s.recvfrom(65535)[0]
        if len(packet) < 42:
            continue
        arp = parse_arp(packet)
        if arp['opcode'] == 1:
            print(
                f"ARP Request: {arp['src_ip']} is asking about {arp['dst_ip']}")
        elif arp['opcode'] == 2:
            print(
                f"ARP Reply: {format_mac(arp['src_mac'])} has address {arp['src_ip']}")


if __name__ == "__main__":
    main()
