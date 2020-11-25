# ![](assets/icon.png) Goholint

![](assets/header.gif)

An incomplete, buggy and suboptimal GameBoy emulator written in Go purely for
fun. It displays pixels, makes noises and can export GIFs. Hopefully it can
prove remotely informative!


## Disclaimer

All code available here is raw, broken, doesn't follow any specific kind of
workflow and isn't guaranteed to work in any way or form.

It might contain blatant errors, awkward workarounds and the occasional
profanity in comments or commit messages. Golang at least guarantees that
the formatting is somewhat consistent.

Those are the main reasons why this code used to be self-hosted for so long.


## Getting Goholint

I started this project before Go modules were a thing so the easiest way to run
Goholint for now is still probably to use `go get`.

```shell
go get github.com/lazy-stripes/goholint
$GOPATH/bin/goholint -help
```


## Usage

The emulator ships without any kind of ROM for hopefully obvious reasons. If
you want a scrolling logo, the emulator needs a boot ROM it will attempt to
read from `bin/boot/dmg_rom.bin` or whatever path you specify with `-boot`.

(Note: if you don't want to hunt down the GameBoy's boot ROM, simply start the
emulator with the `‑fastboot` parameter to bypass it entirely. It doesn't work
as well as using the boot ROM yet, alas.)

With that taken care of, `go run main.go ‑rom <path>` should be enough to see
an SDL window potentially displaying some interesting things, or more likely a
blank screen, if it doesn't crash first.

(As of 2020, Tetris and Dr. Mario are kind of playable!)


## Keys

I plan on making controls configurable, but in the mean time the following keys
are hardcoded into the emulator:

Button          | Key
---             | ---
**A**           | S
**B**           | D
**Select**      | Backspace
**Start**       | Return
**Directions**  | Arrow Keys


(It's sort of okay on QWERTY and AZERTY keyboards alike but *does* make Metroid
II hard to play.)


## Acknowledgements

The present project only exists thanks to Tomek Rękawek and his fascinating
blog article about [how relatively easy it was to start implementing a GB
emulator](https://blog.rekawek.eu/2017/02/09/coffee-gb/). His own emulator,
[coffee-gb](https://github.com/trekawek/coffee-gb), was a great help and
inspiration in the making of this.

I otherwise officially blame @dmuth for retweeting the aforementioned blog post.

Also @balinares for motivating me to make this public. Love you guys! ♥
