# Bit overkill but procrastination knows no bounds.
all: noop.gb

noop.gb: noop.o
	rgblink -o noop.gb noop.o
	rgbfix -f hg noop.gb

noop.o: noop.asm header.asm
	rgbasm -l -o noop.o noop.asm
