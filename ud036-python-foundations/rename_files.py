"""

@author: Phil White

Strips numbers from filenames of target directory
"""

import os

target = os.path.expanduser(r'~/Downloads/prank')
for filename in os.listdir(target):
    os.rename(target + '/' + filename, target + '/' + filename.strip('0123456789'))

