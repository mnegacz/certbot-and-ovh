#/bin/sh
set -e

register() {
  DOMAIN=$1
  echo "Registering $DOMAIN"
  certbot \
    --agree-tos \
    --non-interactive \
    --manual \
    --manual-auth-hook /usr/local/bin/auth \
    --manual-cleanup-hook /usr/local/bin/clean \
    --domain $DOMAIN \
    --domain *.$DOMAIN \
    --email $EMAIL \
    --preferred-challenges dns-01 \
    --server https://acme-v02.api.letsencrypt.org/directory \
    --renew-by-default \
    --expand \
    --manual-public-ip-logging-ok \
    certonly
}

renew() {
  echo "Renewing"
  certbot \
    --manual-auth-hook /usr/local/bin/auth \
    --manual-cleanup-hook /usr/local/bin/clean \
    renew
}

while read DOMAIN; do
  if [ ! -d "/etc/letsencrypt/live/$DOMAIN" ]; then
    register "$DOMAIN"
  fi
done </domains.txt

renew
