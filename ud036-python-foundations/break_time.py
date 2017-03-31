import webbrowser
import time

interval = 0;
while (interval < 3):
    time.sleep(10)
    webbrowser.open('http://www.bbc.co.uk/news')
    interval = interval + 1
