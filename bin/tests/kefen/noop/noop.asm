; Dummy ROM that stops right after BootROM is done executing. Used to debug
; all memory writes occurring during boot process.
INCLUDE "header.asm"

SECTION "default", ROM0
main:
	JR main
