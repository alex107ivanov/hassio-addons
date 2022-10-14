# hassio-addons

## Installation

The installation of this add-ons is pretty simple like installation any other Hass.io add-on.

1. Add this Hass.io add-ons [[repository](https://github.com/alex107ivanov/hassio-addons)] to your Hass.io instance.
2. After some time new addons will appear in list of available addons.
3. Install needed add-on.
4. Start installed add-on.
5. Check logs and status of installed add-on to see if everything went well.

## Addons list

### Watchdog Dev
Plugin activates /dev/watchdog - hardware watchdog device to restart server on no responce. For details about watchdog see https://www.kernel.org/doc/Documentation/watchdog/watchdog-api.txt.
I checked it with my Raspberry Pi 4 - it has Broadcom BCM2835 Watchdog timer, enabled by default.
Service sends keepalive to watchdog timer every 5 seconds, on hang or other software problems system will do hardware restart in 15 seconds.

To get notifications on HA start/stop you can add this automations:

```
- id: "Hass Startup Notification"
  alias: 'Hass Startup Notification'
  trigger:
    - platform: homeassistant
      event: start
  action:
    service: notify.telegram_notify
    data:
      title: "Warning"
      message: "Hass restarted"

- id: "Hass Shutdown Notification"
  alias: 'Hass Shutdown Notification'
  trigger:
    - platform: homeassistant
      event: shutdown
  action:
    service: notify.telegram_notify
    data:
      title: "Warning"
      message: "Hass shutdown"
```

If you received only "restarted" notification - looks like HA restarted uncorrectly.

### OpenVPN Client
Plugin to connect HA to OpenVPN server. Config file must contain all keys.
