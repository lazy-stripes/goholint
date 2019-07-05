# Bit overkill but procrastination knows no bounds.
all: miup.gb

miup.gb: miup.o tiles.bin
	rgblink -o miup.gb miup.o
	rgbfix -f hg miup.gb

miup.o: miup.asm header.asm
	rgbasm -o miup.o miup.asm

tiles.bin: tiles.png
	rgbgfx -o tiles.bin tiles.png
