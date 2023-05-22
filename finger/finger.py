import sys
from cryptography import x509
from cryptography.hazmat.primitives import hashes
from cryptography.hazmat.primitives.serialization import Encoding
from cryptography.hazmat.backends import default_backend
import binascii

# Get the filename from the command line arguments
if len(sys.argv) != 2:
    print("Usage: script.py <filename>")
    sys.exit(1)

filename = sys.argv[1]

# Load the certificate
with open(filename, 'rb') as cert_file:
    cert_data = cert_file.read()
    cert = x509.load_pem_x509_certificate(cert_data, default_backend())

# Compute the SHA256 fingerprint
fingerprint = cert.fingerprint(hashes.SHA256())

# Convert to hex
fingerprint_hex = binascii.hexlify(fingerprint).decode()

# Print the lowercase fingerprint without colons
print(fingerprint_hex.lower())
