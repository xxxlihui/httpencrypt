FROM debian:buster
COPY httpencrypt /
RUN chmod +x /httpencrypt
VOLUME /data
