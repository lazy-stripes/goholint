; Dummy ROM that stops right after BootROM is done executing. Used to debug
; all memory writes occurring during boot process.
INCLUDE "header.asm"

DEF CYCLE_DELAY EQU $04

SECTION "default", ROM0
main:

init_tiles:
	; Turn LCD off to safely access VRAM.
	XOR A
	LD [$FF00+$40],A

	; Default palette: 3210
	LD A, $E4
	LD [$FF00+$48], A

	; Copy tile data to VRAM.
	LD DE, tiles_data
	LD HL, $8000
	LD BC, tiles_data.end - tiles_data

.copy_loop:
	LD A, [DE]
	INC DE
	LD [HLI], A
	DEC BC	; Does not set zero flag so we need to check
	XOR A
	CP B
	JR NZ, .copy_loop
	CP C
	JR NZ, .copy_loop

init_map:
	LD HL, $9800
.copy_loop:
	; Select tile to write. TODO: use RLE-encoded map.
	; X = ADDR / 20, Y = ADDR % 18
	LD A, $01 ; Horizontal tile only for now
.write_tile:
	LD [HLI], A
	LD A, $9c
	CP H
	JR NZ, .copy_loop

	; Turn LCD back on
	LD A, $93
	LD [$FF00+$40],A

	LD E, CYCLE_DELAY

	; Enable vblank interrupt, change palette from there every few frames.
	LD A, $01
	LD [$FF00+$FF], A
	EI

lock:
	JR lock

vblank:
	DEC E
	LD A, E
	JR Z, shiftpal
	RETI

shiftpal:
	; Shift palette bits.
	LD A, [$FF00+$47] ; BGP
	RLC A
	RLC A
	LD [$FF00+$47], A
	LD E, CYCLE_DELAY
	RETI

; Unhandled interrupts.
stat:
timer:
serial:
joypad:
	RETI

; Raw tiles data (3 tiles, 6 bytes)
tiles_data:
	INCBIN "palette-cycle.tiles"
.end

; Tilemap (format is 0bRRRRRIII where R is how many times the tile repeats, and
; I the tile ID).
; 🭽▔🭾 🭼▁🭿 ▏▕
tilemap_data:
	DB $00, $94, $01 		; 🭽▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔🭾
	DB $05, $84, $01, $05 	; ▏🭽▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔🭾▕
	DB $05, $84, $01, $05 	; ▏▏🭽▔▔▔▔▔▔▔▔▔▔▔▔▔▔🭾▕▕
