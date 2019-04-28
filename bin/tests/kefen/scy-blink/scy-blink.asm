; Tiny test to try rgbds. Break scroll and see...

INCLUDE "header.asm"

FRAMES EQU $20

SECTION "default", ROM0
main:
	LD B, FRAMES	; Number of frames between changes

	; Initialize SCY to high position
	LD A, $32
	LD [$FF00+$42], A

	; Enable VBlank interrupt
	LD A, $01
	LD [$FF00+$FF], A
	EI

endless:
	JR endless

; Interrupts
vblank:
	; Wait the required number of frames before changing SCY.
	DEC B
	JR NZ, vblank_done

	; Reset B
	LD B, FRAMES

	; SCY = -SCY (two's complement A)
	LD A, [$FF00+$42]
	CPL
	INC A
	LD [$FF00+$42], A

vblank_done:
	RETI

; Unhandled interrupts.
stat:
timer:
serial:
joypad:
    RETI