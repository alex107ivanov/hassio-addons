FROM alpine:3.7

# Install runtime dependencies. The versions are pinned for reproducible,
# deterministic, pure builds.
RUN apk --update add openvpn jq

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
