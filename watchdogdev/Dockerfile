FROM alpine:3.15

RUN \
      apk update && apk add --no-cache \
      bash \
      py3-pip


RUN \
    echo "unicode=\"YES\"" >> /etc/rc.conf && \
    apk add --no-cache --virtual .build_deps \
        autoconf file g++ gcc libc-dev make pkgconf python3 python3-dev linux-headers

RUN \
      pip3 install watchdogdev

COPY watchdog.py /
RUN chmod +x /watchdog.py

CMD ["/usr/bin/python3", "/watchdog.py"]
