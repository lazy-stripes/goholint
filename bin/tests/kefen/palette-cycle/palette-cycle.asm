; Dummy ROM that stops right after BootROM is done executing. Used to debug
; all memory writes occurring during boot process.
INCLUDE "header.asm"
INCLUDE "../include/macros.asm"

DEF CYCLE_DELAY EQU $08

SECTION "default", ROM0
main:

init_tiles:
	; Turn LCD off to safely access VRAM.
	XOR A
	LD [$FF00+$40],A

	; Default palette: 3 2 1 0
	LD A, $E4
	LD [$FF00+$47], A

	; Reset SCX that the Goholint boot ROM set.
	; FIXME: Goholint shouldn't do that, rework the logo instead.
	XOR A
	LD [$FF00+$43], A

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
.init_loop:
	LD C, $14	; Init C to 20 (width in tiles)
.copy_loop:
	; Read an entry from the tilemap data and write as many needed IDs
	CALL write_tiles

	; Check if C is zero, if so increment HL offset to next visible line
	XOR A
	OR C
	JR NZ, .copy_loop


	; Check if we're done.
	LD A, E
	;OutputA

	CP LOW(tilemap_data.end)


	JR Z, init_end

	; We got to tile 20, skip the next 12 bytes.
	LD BC, $000c
	ADD HL, BC
	JR .init_loop

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
	DEC C
	DEC B
	JR NZ, .copy	; Keep copying while count is nonzero.
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

; Align tile data to byte so we can just check E to know if we finished reading
; all tiles.
SECTION "tilemap_data", ROM0, ALIGN[8]

; Tilemap (format is 0bRRRRRIII where R is how many times the tile repeats, and
; I the tile ID). Count (R) is 1-based.
; /=0 \=1 \=2 /=3 ⁻=4 _=5 |<=6 >|=7
tilemap_data:
	DB      $08, $94, $09 		; +------------------+
	DB $0e, $08, $84, $09, $0f 	; |+----------------+|
	DB $16, $08, $74, $09, $17 	; ||+--------------+|| 0001 0 110
	DB $1e, $08, $64, $09, $1f 	; |||+------------+|||
	DB $26, $08, $54, $09, $27 	; ||||+----------+||||
	DB $2e, $08, $44, $09, $2f 	; |||||+--------+|||||
	DB $36, $08, $34, $09, $37 	; ||||||+------+||||||
	DB $3e, $08, $24, $09, $3f 	; |||||||+----+|||||||
	DB $46, $08, $14, $09, $47 	; ||||||||+--+||||||||
	DB $46, $0a, $15, $0b, $47 	; ||||||||+--+||||||||
	DB $3e, $0a, $25, $0b, $3f 	; |||||||+----+|||||||
	DB $36, $0a, $35, $0b, $37 	; ||||||+------+||||||
	DB $2e, $0a, $45, $0b, $2f 	; |||||+--------+|||||
	DB $26, $0a, $55, $0b, $27 	; ||||+----------+||||
	DB $1e, $0a, $65, $0b, $1f 	; |||+------------+|||
	DB $16, $0a, $75, $0b, $17 	; ||+--------------+||
	DB $0e, $0a, $85, $0b, $0f 	; |+----------------+|
	DB      $0a, $95, $0b 		; +------------------+
.end
