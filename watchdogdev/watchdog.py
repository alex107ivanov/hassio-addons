#!/usr/bin/env python3
from time import sleep
from watchdogdev import *
import signal
import logging

class WatchdogDev:
  def __init__(self):
    self.shutdown = False
    signal.signal(signal.SIGINT, self.exit_gracefully)
    signal.signal(signal.SIGTERM, self.exit_gracefully)
    self.sleep_time = 5

  def exit_gracefully(self, signum, frame):
    logging.info('Received signal: %d', signum)
    self.shutdown = True

  def run(self):
    logging.info("Opening watchdog device")
    self.wdt = watchdog('/dev/watchdog')
    self.wdt.get_support()
    logging.info("Watchdog identity: %s", self.wdt.identity)
    logging.info("Watchdog firmware version: %d", self.wdt.firmware_version)
    logging.info("Watchdog options: %d", self.wdt.options)
    logging.info("Watchdog timeout: %d", self.wdt.get_timeout())
    logging.info("Starting main cycle with sleep time %d sec...", self.sleep_time)
    while self.shutdown != True:
      logging.debug("Watchdog timeout left: %d", self.wdt.get_time_left())
      logging.debug("Send keep alive")
      self.wdt.keep_alive()
      sleep(self.sleep_time)
    logging.info("Got shutdown signal")
    if self.wdt.options & WDIOF_MAGICCLOSE == WDIOF_MAGICCLOSE:
      logging.info("Magic close needed")
      self.wdt.magic_close()
      logging.info("Magic close done")
    logging.info("Main cycle finished")

if __name__ == '__main__':
  logging.basicConfig(format='%(asctime)s %(levelname)-8s %(message)s', level=logging.INFO, datefmt='%Y-%m-%d %H:%M:%S')
  app = WatchdogDev()
  app.run()
  logging.info("Exit")
