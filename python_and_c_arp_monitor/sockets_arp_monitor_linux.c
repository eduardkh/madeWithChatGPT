#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <net/if_arp.h>      // Use this instead of netinet/if_arp.h
#include <netinet/ether.h>   // For Ethernet header
#include <arpa/inet.h>
#include <unistd.h>

#define BUFSIZE 65536

int main() {
    int sockfd;
    unsigned char buffer[BUFSIZE];

    sockfd = socket(AF_PACKET, SOCK_RAW, htons(ETH_P_ARP));
    if (sockfd < 0) {
        perror("Socket error");
        exit(EXIT_FAILURE);
    }

    printf("Listening for ARP packets...\n");

    while (1) {
        ssize_t num_bytes = recvfrom(sockfd, buffer, BUFSIZE, 0, NULL, NULL);
        if (num_bytes < 0) {
            perror("Recv error");
            continue;
        }

        struct ether_header *eth = (struct ether_header *) buffer;
        if (ntohs(eth->ether_type) == ETH_P_ARP) {
            struct ether_arp *arp = (struct ether_arp *)(buffer + sizeof(struct ether_header));

            // Determine if it's a request or a reply.
            int op = ntohs(arp->ea_hdr.ar_op);
            char *op_str;
            if (op == ARPOP_REQUEST)
                op_str = "Request";
            else if (op == ARPOP_REPLY)
                op_str = "Reply";
            else
                op_str = "Unknown";

            printf("ARP %s: %zd bytes\n", op_str, num_bytes);
            printf("Sender MAC: %02x:%02x:%02x:%02x:%02x:%02x\n",
                arp->arp_sha[0], arp->arp_sha[1], arp->arp_sha[2],
                arp->arp_sha[3], arp->arp_sha[4], arp->arp_sha[5]);
            printf("Sender IP: %d.%d.%d.%d\n",
                arp->arp_spa[0], arp->arp_spa[1], arp->arp_spa[2], arp->arp_spa[3]);
            printf("Target IP: %d.%d.%d.%d\n",
                arp->arp_tpa[0], arp->arp_tpa[1], arp->arp_tpa[2], arp->arp_tpa[3]);
        }
    }

    close(sockfd);
    return 0;
}
