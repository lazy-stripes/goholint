; Dummy ROM that stops right after BootROM is done executing. Used to debug
; all memory writes occurring during boot process.
INCLUDE "header.asm"

SECTION "default", ROM0
main:
	; Enable vblank interrupt, and disable display from there.
	LD A, $01
	LD [$FF00+$FF], A
	EI

lock:
	JR lock

vblank:
	XOR A
	LD [$FF00+$40], A ; Turn off display
	RET

; Unhandled interrupts.
stat:
timer:
serial:
joypad:
	RETI