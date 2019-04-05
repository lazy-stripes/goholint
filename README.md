# gb.go

An incomplete, buggy and suboptimal GameBoy emulator written in Go purely for
fun.


# Disclaimer

All code available here is raw, broken, doesn't follow any specific kind of
workflow and isn't guaranteed to work in any way or form.

It might contain blatant errors, awkward workarounds and the occasional
profanity in comments or commit messages. Golang at least guarantees that
the formatting is somewhat consistent.

Those are the main reasons why this is self-hosted and not available from a
serious platform like GitHub yet.


# Installation

## Using `go get`

This is the easiest method and should (hopefully) work out of the box.

```shell
go get go.tigris.fr/gameboy
```

## Using `git.tigris.fr`

Alternatively, directly using `git` (which is pretty much what `go get` does)
should work too. As of this writing, `tigris.fr` only offers read-only access to
public repositories.

```shell
cd $GOPATH/src
mkdir -p go.tigris.fr
cd go.tigris.fr
git clone https://git.tigris.fr/public/gameboy.git
```


# Usage

The emulator ships without any kind of ROM for hopefully obvious reasons. If
you want a scrolling logo, the emulator needs a boot ROM it will attempt to
read from `bin/DMG_ROM.bin`.

(Note: if you don't want to hunt down the GameBoy's boot ROM, simply start the
emulator with the `‑fastboot` parameter to bypass it entirely.)

With that taken care of, `go run main.go ‑rom <path>` should be enough to see
an SDL window potentially displaying some interesting things, or more likely a
blank screen, if it doesn't crash first.


# Acknowledgements

The present project only exists thanks to Tomek Rękawek and his fascinating
blog article about [how relatively easy it was to start implementing a GB
emulator](https://blog.rekawek.eu/2017/02/09/coffee-gb/).
As such, large chunks of the present source code are heavily inspired by
[coffee-gb's source code](https://github.com/trekawek/coffee-gb).

I otherwise officially blame @dmuth for retweeting the aforementioned blog post.

Also @balinares for motivating me to make this public. Love you guys! ♥
