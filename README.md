# hassio-addons

## Watchdog Dev
Plugin activates /dev/watchdog - hardware watchdog device to restart server on no responce. For details about watchdog see https://www.kernel.org/doc/Documentation/watchdog/watchdog-api.txt.
I checked it with my Raspberry Pi 4 - it has Broadcom BCM2835 Watchdog timer, enabled by default.
Service sends keepalive to watchdog timer every 5 seconds, on hang or other software problems system will do hardware restart in 15 seconds.
