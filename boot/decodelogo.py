#!/usr/bin/python3

# Quick script to turn the given PNG into a list of DB instructions in GameBoy
# assembly. For use in our custom boot ROM's assembly source.

import sys

from PIL import Image

# Create Image object
logo = Image.open('logo.png')

if logo.width != 48 or logo.height != 8:
    print("Logo must be 48x8 black and white.")
    sys.exit(1)

values = []

# Decode 2 rows of 4×4 pixels to make bytes.
for y in range(0, logo.height, 4):
    for x in range(0, logo.width, 4):
        v = 0x00
        for offset in range(16):
            pix = logo.getpixel((x+(offset%4), y+(offset/4)))
            bit = int(pix[0] == 0)
            v |= bit

            if offset == 7 or offset == 15:
                values.append(v)
                v = 0x00
            else:
                v <<= 1

print("DB %s" % (",".join("$%02X" % v for v in values[0:16])))
print("DB %s" % (",".join("$%02X" % v for v in values[16:32])))
print("DB %s" % (",".join("$%02X" % v for v in values[32:48])))
