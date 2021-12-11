#!/usr/bin/env python3
from time import sleep
from watchdogdev import *
import signal

class WatchdogDev:
  def __init__(self, timeout):
    self.shutdown = False
    signal.signal(signal.SIGINT, self.exit_gracefully)
    signal.signal(signal.SIGTERM, self.exit_gracefully)
    self.timeout = timeout
    if self.timeout < 10:
      self.timeout = 10
    self.sleep_time = self.timeout / 10
    if self.sleep_time < 1:
      self.sleep_time = 1

  def exit_gracefully(self, signum, frame):
    print('Received signal:', signum)
    self.shutdown = True

  def run(self):
    print("Opening watchdog device")
    self.wdt = watchdog('/dev/watchdog')
    self.wdt.get_support()
    print("Watchdog identity: ", self.wdt.identity)
    print("Watchdog firmware version: ", self.wdt.firmware_version)
    print("Watchdog options: ", self.wdt.options)
    print("Current timeout: ", self.wdt.get_timeout())
    self.wdt.set_timeout(timeout)
    print("New timeout: ", self.wdt.get_timeout())
    print("Starting main cycle...")
    while self.shutdown != True:
      print("Timeout left: ", self.wdt.get_time_left())
      print("Send Keep alive")
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
  app = WatchdogDev(60)
  app.run()
  print("Exit")
