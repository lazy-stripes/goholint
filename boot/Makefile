all: boot.rom

install: boot.rom
	cp -f boot.rom ../assets/boot.rom

clean:
	rm -f boot.o boot.rom

boot.rom: boot.o
	rgblink -x -o boot.rom boot.o

boot.o: boot.asm
	rgbasm -l -o boot.o boot.asm