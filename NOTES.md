# Development notes

The current goal of this branch is to move the UI out of the Gameboy code.

At the moment it feels that the following might be achieved relatively easily:

* Move all UI calls made from Gameboy code to the UI package itself:
  * Messages
  * Screenshots
  * GIFs
* Have a Gameboy and UI instance exist in parallel.
* Repaint the UI when the Gameboy returns vblank=true after a tick.
* ...
* Profit?
