#!/usr/bin/env python3
from time import sleep
from watchdogdev import *

wdt = watchdog('/dev/watchdog')
wdt.get_support()
print("Watchdog identity: ", wdt.identity)
print("Watchdog firmware version: ", wdt.firmware_version)
print("Watchdog options: ", wdt.options)
print("Current timeout: %d" % wdt.get_timeout())

wdt.set_timeout(60)
print("New timeout: %d" % wdt.get_timeout())

for i in range(5):
    print("Send Keep alive %d" % i)
    wdt.keep_alive()
    for j in range(5): #Change to 15 to see RPi reboot ...
        print("... Waiting", j, ", Left :", wdt.get_time_left())
        sleep(1)

if wdt.options & WDIOF_MAGICCLOSE == WDIOF_MAGICCLOSE:
  print("Magic Close Needed")
  wdt.magic_close()

print("Done !")