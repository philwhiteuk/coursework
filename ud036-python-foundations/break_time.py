"""

@author: Phil White

Background script to remind the user to take a break from the screen every 2 hours.
Runs on weekdays, between the hours of 9-5
"""

import webbrowser
import time
import datetime
date = datetime.date

while True:
    time.sleep(60 * 60 * 2)
    today = date.today()

    if today.isoweekday() < 6 and 9 <= time.localtime().tm_hour <= 17:
        webbrowser.open('http://www.bbc.co.uk/news')

