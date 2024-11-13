all: palette-cycle.gb

clean:
	rm -f palette-cycle.o palette-cycle.tiles palette-cycle.gb

palette-cycle.gb: palette-cycle.o
	rgblink -o palette-cycle.gb palette-cycle.o
	rgbfix -f hg -p 255 palette-cycle.gb

palette-cycle.o: palette-cycle.asm header.asm palette-cycle.tiles
	rgbasm -o palette-cycle.o palette-cycle.asm

palette-cycle.tiles: palette-cycle-tiles.png
	rgbgfx -o palette-cycle.tiles palette-cycle-tiles.png

