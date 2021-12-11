#!/usr/bin/env python3
from time import sleep
from watchdogdev import *
import signal

class WatchdogDev:
  def __init__(self):
    self.shutdown = False
    signal.signal(signal.SIGINT, self.exit_gracefully)
    signal.signal(signal.SIGTERM, self.exit_gracefully)
    self.sleep_time = 5

  def exit_gracefully(self, signum, frame):
    print('Received signal:', signum)
    self.shutdown = True

  def run(self):
    print("Opening watchdog device")
    self.wdt = watchdog('/dev/watchdog')
    self.wdt.get_support()
    print("Watchdog identity:", self.wdt.identity)
    print("Watchdog firmware version:", self.wdt.firmware_version)
    print("Watchdog options:", self.wdt.options)
    print("Current timeout:", self.wdt.get_timeout())
    print("Starting main cycle with sleep time", self.sleep_time,"sec...")
    while self.shutdown != True:
      print("Watchdog timeout left:", self.wdt.get_time_left())
      print("Send keep alive")
      self.wdt.keep_alive()
      sleep(self.sleep_time)
    print("Got shutdown signal")
    if wdt.options & WDIOF_MAGICCLOSE == WDIOF_MAGICCLOSE:
      print("Magic close needed")
      self.wdt.magic_close()
      print("Magic close done")
    print("Main cycle finished")

  def stop(self):
    print("Stop app")

if __name__ == '__main__':
  app = WatchdogDev()
  app.run()
  print("Exit")
