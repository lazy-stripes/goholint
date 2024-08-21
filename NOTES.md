# Development notes

The current goal of this branch is to move the UI out of the Gameboy code.

At the moment it feels that the following might be achieved relatively easily:

* Move all UI calls made from Gameboy code to the UI package itself:
  * Messages
  * Screenshots
  * GIFs
* Put UI first (branch name should be a hint) and do all high-level stuff
  (screenshots, GIF, etc) from it.
* Repaint the UI when the Gameboy returns vblank=true after a tick.
* ...
* Profit?
