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
	LD DE, tilemap_data
	LD HL, $9800
.copy_loop:
	; Read an entry from the tilemap data and write as many needed IDs
	CALL write_tiles

	; TODO: jump to next visible line when we reach 20 tiles. Use C for counting?


write_tiles:
	LD A, [DE]	; Read next entry
	INC DE

	LD B, A 	; Copy A to keep the count bits
	SRL B		; Extract top five bits (count) by shifting right ×3.
	SRL B
	SRL B

	AND A, $07	; Mask count bits to keep tile ID

.copy:
	LD [HLI], A
	DEC B
	JR NZ, .copy
	RET

init_end:
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
; I the tile ID). TODO: count (R) should be 1-based. Add 4 to all values.
; /=0 \=1 \=2 /=3 -=4 |=5
tilemap_data:
	DB $00, $94, $01 				; +------------------+
	DB $05, $00, $84, $01, $05 		; |+----------------+|
	DB $15, $00, $74, $01, $15 		; ||+--------------+||
	DB $1d, $00, $64, $01, $1d 		; |||+------------+|||
	DB $25, $00, $54, $01, $25 		; ||||+----------+||||
	DB $2d, $00, $44, $01, $2d 		; |||||+--------+|||||
	DB $35, $00, $34, $01, $35 		; ||||||+------+||||||
	DB $3d, $00, $24, $01, $3d 		; |||||||+----+|||||||
	DB $45, $00, $14, $01, $45 		; ||||||||+--+||||||||
	DB $45, $02, $14, $03, $45 		; ||||||||+--+||||||||
	DB $3d, $02, $24, $03, $3d 		; |||||||+----+|||||||
	DB $35, $02, $34, $03, $35 		; ||||||+------+||||||
	DB $2d, $02, $44, $03, $2d 		; |||||+--------+|||||
	DB $25, $02, $54, $03, $25 		; ||||+----------+||||
	DB $1d, $02, $64, $03, $1d 		; |||+------------+|||
	DB $15, $02, $74, $03, $15 		; ||+--------------+||
	DB $05, $02, $84, $03, $05 		; |+----------------+|
	DB $02, $94, $03 				; +------------------+
.end