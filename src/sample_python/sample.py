#!/usr/bin/python3
 
import os
 
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
 
CHROME_BIN = "/usr/bin/chromium-browser"
CHROME_DRIVER = os.path.expanduser('/usr/bin/chromedriver')
 
options = Options()
options.binary_location = CHROME_BIN
options.add_argument('--headless')
options.add_argument('--window-size=1280,3000')
 
driver = webdriver.Chrome(CHROME_DRIVER, chrome_options=options)
 
driver.get("https://dev.classmethod.jp")
driver.save_screenshot("./devio.png")
driver.quit()
