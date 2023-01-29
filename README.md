# PNGme
Utility for storing text messages in a png file.

## Usage examples:

### Encoding message in png file:
``` ./pngme encode -file=./robocop.png -head "teST" -msg "some message" ```

### Decoding message from png file:
``` ./pngme decode -file=./robocop.png -head "teST" ```

### Remove message from png file:
``` ./pngme remove -file=./robocop.png -head "teST" ```

### Print all chunks in png file: 
``` ./pngme print -file=./robocop.png```