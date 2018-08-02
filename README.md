# gb.go

An incomplete, buggy and suboptimal GameBoy emulator written in Go purely for fun.


# Disclaimer

All code available here is raw, broken, doesn't follow any specific kind of workflow and isn't guaranteed to work in
any way or form.

It might contain blatant errors, awkward workarounds and the occasional profanity in comments or commit messages.
Golang at least guarantees that the formatting is somewhat consistent.

Those are the main reasons why this is self-hosted and not available from a serious platform like GitHub yet.


# Installation

## Using go get

This is the easiest method and should (hopefully) work out of the box.

```shell
go get go.tigris.fr/gameboy
```

## Using git.tigris.fr

Alternatively, directly using git (which is pretty much what `go get` does) should work too.
As of this writing, tigris.fr only offers read-only access to public repositories.

```shell
cd $GOPATH/src
mkdir -p go.tigris.fr
cd go.tigris.fr
git clone https://git.tigris.fr/gameboy.git
```


# Usage

The emulator ships without any kind of ROM for hopefully obvious reasons. It expects them to be in the `bin/`
folder.

With that taken care of, `go run main.go` should be enough to see an SDL window showing a disabled GameBoy screen for
a long time, and possibly some interesting things for a brief couple seconds.


# Acknowledgements

The present project only exists thanks to Tomek Rękawek and his fascinating blog article about
[how relatively easy it was to start implementing a GB emulator](https://blog.rekawek.eu/2017/02/09/coffee-gb/).
As such, large chunks of the present source code are heavily inspired by
[coffee-gb's source code](https://github.com/trekawek/coffee-gb).

I otherwise officially blame @dmuth for retweeting the aforementioned blog post.

Also @balinares for motivating me to make this public. Love you guys! ♥
