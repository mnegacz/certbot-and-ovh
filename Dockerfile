FROM certbot/certbot

VOLUME /etc/letsencrypt

ENV DOMAIN example.org
ENV EMAIL example@example.org

COPY auth /usr/local/bin/
COPY clean /usr/local/bin/
COPY ovh.conf /etc/
COPY entrypoint.sh /

ENTRYPOINT /entrypoint.sh
