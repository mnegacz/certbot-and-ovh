FROM certbot/certbot

VOLUME /etc/letsencrypt

ENV EMAIL example@example.org

COPY auth /usr/local/bin/
COPY clean /usr/local/bin/
COPY ovh.conf /etc/
COPY domains.txt /
COPY entrypoint.sh /

ENTRYPOINT /entrypoint.sh
