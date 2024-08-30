; Custom boot ROM code that uses the Goholint logo and a slightly faster boot
; sequence. Also we don't set memory to zero here, the emulator already does it.
SECTION "default", ROM0[$0000]
main:
	LD SP, $fffe		; Setup Stack

	LD HL, $ff26		; Setup Audio
	LD C, $11
	LD A, $80
	LD [HLD], A			; $ff26 = $80 → NR52 bit 7 = 1 → All sound on
	LD [$FF00+C], A		; $ff11 = $80 → NR11 bit 7-6 = 10 → Wave Duty 50%
	INC C
	LD A, $f3
	LD [$FF00+C], A		; $ff12 = $f3 → Initial volume $0f, Decrease, Sweep 3
	LD [HLD], A			; $ff25 = $f3 → Output sounds 1-4 to S02, 1-2 to S01
	LD A, $77
	LD [HL], A			; $ff24 = $77 → S02-S01 volume 7

	LD A, $04			; Setup SCX
	LD [$FF00+$43], A

	LD A, $9c			; Setup BG palette
	LD [$FF00+$47], A

	LD DE, logo_data	; Convert and load logo data from ROM into Video RAM
	LD HL, $8010
.logo_loop:
	LD A, [DE]
	CALL scale_logo_1
	CALL scale_logo_2
	INC DE
	LD A, E
	CP LOW(acorn_tiles)	; Did we reach the end of logo (start of tile data)?
	JR NZ, .logo_loop

	LD DE, acorn_tiles	; Load acorn tiles into Video RAM.
	LD B, $20
.tiles_loop:
	LD A, [DE]
	INC DE
	LD [HLI], A
	DEC B
	JR NZ, .tiles_loop

	LD A, $1a			; Setup background tilemap
	LD [$9910], A		; Acorn (top tile)
	DEC A
	LD [$9930], A		; Acorn (bottom tile)

	LD HL, $992f
.tilemap_init:
	LD C, $0c
.tilemap_loop:
	DEC A
	JR Z, scroll_logo
	LD [HLD], A
	DEC C
	JR NZ, .tilemap_loop
	LD L, $0f
	JR .tilemap_init

; Scrolling boot sequence
scroll_logo:
	LD H, A				; Initialize scroll count, H=0
	LD A, $64
	LD D, A				; Set loop count, D=$64
	LD [$FF00+$42], A	; Set vertical scroll register
	LD A, $91
	LD [$FF00+$40], A	; Turn on LCD, showing Background
	INC B				; Set B=1
.start_scroll:
	LD E, $02
.init_wait_frame:
	LD C, $08			; Number of frames to wait before scrolling more.
.wait_frame:
	LD A, [$FF00+$44]	; Wait for screen frame, repeat until C = 0.
	CP $90
	JR NZ, .wait_frame
	DEC C
	JR NZ, .wait_frame
	DEC E
	JR NZ, .init_wait_frame

	LD C, $13
	INC H				; Increment scroll count
	LD A, H
	LD E, $83
	CP $62				; $62 counts in, play sound #1
	JR Z, .beep
	LD E, $c1
	CP $64
	JR NZ, .scroll		; $64 counts in, play sound #2
.beep:
	LD A, E				; Play sound
	LD [$FF00+C], A

	INC C
	LD A, $87
	LD [$FF00+C], A

.scroll:
	LD A, [$FF00+$42]
	SUB B
	LD [$FF00+$42], A	; Scroll logo up if B=1
	DEC D
	JR NZ, .start_scroll

	DEC B				; Set B=0 first time
	JR NZ, check		; ... next time, will jump to end of boot.

	LD D, $20
	JR .start_scroll

; Loading and scaling up raw logo.
scale_logo_1:
	LD C, A				; Scale up all the bits of the graphics data.
scale_logo_2:
	LD B, $04
.scale_loop:
	PUSH BC
	RL C
	RLA
	POP BC
	RL C
	RLA
	DEC B
	JR NZ, .scale_loop
	LD [HLI], A
	INC HL
	LD [HLI], A
	INC HL
	RET

logo_data:
	; Logo (raw bitmap).
	DB $36,$CC,$C6,$00,$00,$07,$01,$19,$08,$8B,$00,$01,$00,$0E,$66,$66
	DB $CC,$0D,$00,$0B,$03,$73,$00,$80,$CC,$E7,$E6,$6C,$CC,$C7,$DD,$D9
	DB $D9,$99,$BB,$B9,$33,$3E,$66,$66,$DD,$DD,$D9,$99,$BB,$BB,$00,$00

acorn_tiles:
	; Top and bottom tiles for the acorn.
	DB $87,$7e,$87,$5e,$87,$5e,$87,$7e,$46,$3c,$4e,$3c,$3c,$18,$18,$00
	DB $00,$00,$18,$00,$24,$18,$3c,$00,$7e,$3c,$df,$7e,$ff,$7e,$ff,$00

check:
	; Check header to lock up if no cartridge is present. I don't have room for
	; much code here, so we just check a header byte for which $ff isn't a valid
	; value.
	LD A, [$014a]		; Destination code, should be $00 or $01.
	CP A, $ff
.check_failure
	JR Z, .check_failure

SECTION "end", ROM0[$00fe]
end:
	; Disable boot ROM before handing over control to the cartridge.
	LD [$FF00+$50], A	; Turn off DMG rom
