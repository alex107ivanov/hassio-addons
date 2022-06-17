#!/bin/sh

set -e

# prepare vpn
mkdir -p /dev/net
if [ ! -c /dev/net/tun ]; then
	mknod /dev/net/tun c 10 200
fi

OVPNCONF=${OVPNCONF:-/vpn/config/config.ovpn}

cat /data/options.json | jq -r '.config' > ${OVPNCONF}

[ -f $OVPNCONF ] || {
	echo $OVPNCONF is missing please provide the volume with the file  > /dev/stderr;
	exit 1;
}

exec /usr/sbin/openvpn --config $OVPNCONF

