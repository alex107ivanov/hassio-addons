ARG BUILD_FROM
FROM $BUILD_FROM

ENV LANG C.UTF-8

RUN apk add --no-cache jq curl file bash openssl

# Copy data for add-on
COPY run.sh /
COPY upload.sh /opt/upload.sh
COPY supload.sh /opt/supload.sh

RUN chmod a+x /run.sh /opt/upload.sh /opt/supload.sh

CMD [ "/run.sh" ]