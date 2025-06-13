#!/usr/bin/env sh
# detect_ca_bundle.sh - Detect default CA bundle location on various Linux distros
# Usage: ./detect_ca_bundle.sh

# Check common CA bundle paths
for path in \
  /etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem \
  /etc/ssl/certs/ca-certificates.crt \
  /etc/ssl/cert.pem \
  /usr/local/share/ca-certificates/ca-certificates.crt; do
  if [ -f "$path" ]; then
    echo "$path"
    exit 0
  fi
done

# Fallback: parse /etc/os-release to guess distribution
if [ -f /etc/os-release ]; then
  . /etc/os-release
  case "$ID_LIKE" in
    *rhel*|*fedora*)
      echo "/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem"
      exit 0
      ;;
    *debian*|*ubuntu*)
      echo "/etc/ssl/certs/ca-certificates.crt"
      exit 0
      ;;
  esac
fi

echo "No system CA bundle found" >&2
exit 1
