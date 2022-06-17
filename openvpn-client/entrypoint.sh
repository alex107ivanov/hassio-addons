#!/bin/sh

set -e

# prepare vpn
mkdir -p /dev/net
if [ ! -c /dev/net/tun ]; then
	mknod /dev/net/tun c 10 200
fi

OVPNCONF='/config.ovpn'

cat /data/options.json | jq --raw-output '.config[]' > $OVPNCONF

exec /usr/sbin/openvpn --config $OVPNCONF

