from scapy.all import *


def arp_display(pkt):
    if pkt.haslayer(ARP):
        if pkt[ARP].op == 1:  # who-has (request)
            return f"ARP Request: {pkt[ARP].psrc} is asking about {pkt[ARP].pdst}"
        elif pkt[ARP].op == 2:  # is-at (response)
            return f"ARP Reply: {pkt[ARP].hwsrc} has address {pkt[ARP].psrc}"


sniff(filter="arp", prn=arp_display, store=0)
