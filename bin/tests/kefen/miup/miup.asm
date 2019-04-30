; Small test to display and scroll sprites, then do the same with a window,
; then question life choices. Still, a good way to test features currently
; missing from the emulator.

INCLUDE "header.asm"

; Constants.
WORK_OAM_ADDR EQU $C000

FADE_DELAY EQU $20

; Utility routine to emulate CALL HL. The idea is to use RST $00 with HL
; containing the address to call. This transparently handles pushing the
; proper return address to the stack.
SECTION "call_hl", ROM0[$0000]
	JP HL
	; The next RET instruction will return right after the RST call.


SECTION "default", ROM0
main:
init_oam:
	; Initialize shadow OAM
	LD DE, sprites_data
	LD HL, WORK_OAM_ADDR
	LD C, sprites_data.end - sprites_data
.copy_loop
	LD A, [DE]
	INC DE
	LD [HLI], A
	DEC C
	JR NZ, .copy_loop

init_fade:
	; Logo doesn't care about the palette's high nibble, but we do: we want
	; to simply substract 4 to the palette byte until it's zero.
	LD A, $0C
	LDH [$FF00+$47], A

init_tiles:
	; Turn LCD off to safely access VRAM. FIXME: gb.go doesn't care. It should!
	XOR A
	LD [$FF00+$40],A

	; Default palette: 3210
	LD A, $E4
	LD [$FF00+$48], A	; FIXME: handle OBJ palettes in gb.go

	; Copy tile data to VRAM.
	LD DE, tiles_data
	LD HL, $81A0	; Right after Nintendo logo tiles
	LD BC, tiles_data.end - tiles_data

.copy_loop
	LD A, [DE]
	INC DE
	LD [HLI], A
	DEC BC	; Does not set zero flag so we need to check
	XOR A
	CP B
	JR NZ, .copy_loop
	CP C
	JR NZ, .copy_loop

	; Turn LCD back on
	LD A, $93
	LD [$FF00+$40],A

init_dma:
	CALL install_dma_routine

start:
	; Start our state machine of sorts: alternate wait with fade, then scroll
	; sprite in.

	LD HL, state_wait_frames	; State to execute at vblank.
	LD C, FADE_DELAY			; Number of frames to wait.
	LD DE, state_fade_logo		; State to execute when delay is elapsed.

	; Enable vblank interrupt, we'll do it all from there.
	LD A, $01
	LD [$FF00+$FF], A
	EI

game_loop:
	HALT	; Just wait for interrupt
	JR game_loop


;
; Interrupts
;

; Call game state defined in HL
vblank:
	; Emulating a CALL HL of sorts. I'm probably doing this wrong.
	RST $00
	RETI

;
; Unhandled interrupts.
;
stat:
timer:
serial:
joypad:
    RETI

;
; Game states
;

; Wait a given number of screen frames.
; Initialize C with the number of vblanks to wait.
; Initialize DE with the address of the next state to execute.
state_wait_frames:
	DEC C
	JR NZ, .end

	; Delay's elapsed, load HL with next state to execute.
	PUSH DE
	POP HL

.end:
	RET


; Make Nintendo logo fade to white. Eebildz sprite is too big to scroll in
; without messing with it anyway.
; BIOS palette is $FC, but we only care about bits 2-3. They'll go from 3 to 0.
state_fade_logo:
	; Reset delay.
	LD C, FADE_DELAY
	LD HL, state_wait_frames

	; Decrease palette entry.
	LDH A, [$FF00+$47]
	SUB 4
	LDH [$FF00+$47], A
	JR NZ, .end

	; Zero reached. Update next state to load after the next delay.
	LD DE, state_update_sprites

.end
	RET

LEAN_IN_DELAY EQU $04	; Sprite scroll speed
state_update_sprites:
	; Decrease X for each sprite until all is visible.
	LD H, HIGH(WORK_OAM_ADDR)
	LD A, $01	; Offset into current OAM entry (X position)
.update_loop:
	LD L, A
	DEC [HL]
	ADD A, $04	; Next entry
	CP A, $A1
	JR NZ, .update_loop

	CALL $FF80	; DMA

	; We're done when the last tile's X position is 'not quite visible'
	LD A, 157
	CP A, [HL]
	JR NZ, .end

	; We're... done for now.
.hang:
	JR .hang

.end
	; Reset delay.
	LD C, LEAN_IN_DELAY
	LD HL, state_wait_frames
	RET


; DMA routine copying data from $C000.
dma_routine:
	LD A, $C0
	LDH [$FF00+$46], A
	LD A, $28
.wait:
	DEC A
	JR NZ, .wait
	RET
.end	; Used to compute routine size


; Copy DMA routine defined above to high RAM ($FF80).
install_dma_routine:
	LD C, $80
	LD B, dma_routine.end - dma_routine
	LD HL, dma_routine
.copy_loop:
	LD A, [HLI]
	LD [$FF00+C], A
	INC C
	DEC B
	JR NZ, .copy_loop
	RET


sprites_data:
	; Alas, large MIUP sprite doesn't quite fit in 40 tiles so I'll cheat
	; and put it in the corner to crop it from the bottom too.
	; It will probably work better to scroll in a window with those tiles.
	; Erm... actually that sprite is way too big, should have done the math.
	; XXX: base coordinates are top-left for now.
	; Using $1A id offset to put our tiles right after the Nintendo logo.
	;   Y    X    #    Palette
	DB $10+$28, $08+$A0, $00+$1A, $00	; #0  ( 8,16)
	DB $10+$28, $10+$A0, $01+$1A, $00	; #1  (16,16)
	DB $10+$28, $18+$A0, $02+$1A, $00	; #2  (24,16)
	DB $10+$28, $20+$A0, $03+$1A, $00	; #3  (32,16)
	DB $10+$28, $28+$A0, $04+$1A, $00	; #4  (40,16)
	DB $18+$28, $08+$A0, $05+$1A, $00	; #5  ( 8,24)
	DB $18+$28, $10+$A0, $06+$1A, $00	; #6  (16,24)
	DB $18+$28, $18+$A0, $07+$1A, $00	; #7  (24,24)
	DB $18+$28, $20+$A0, $08+$1A, $00	; #8  (32,24)
	DB $18+$28, $28+$A0, $09+$1A, $00	; #9  (40,24)
	DB $20+$28, $08+$A0, $0A+$1A, $00	; #10 ( 8,32)
	DB $20+$28, $10+$A0, $0B+$1A, $00	; #11 (16,32)
	DB $20+$28, $18+$A0, $0C+$1A, $00	; #12 (24,32)
	DB $20+$28, $20+$A0, $0D+$1A, $00	; #13 (32,32)
	DB $20+$28, $28+$A0, $0E+$1A, $00	; #14 (40,32)
	DB $20+$28, $30+$A0, $0F+$1A, $00	; #15 (48,32)
	DB $28+$28, $08+$A0, $10+$1A, $00	; #16 ( 8,40)
	DB $28+$28, $10+$A0, $11+$1A, $00	; #17 (16,40)
	DB $28+$28, $18+$A0, $12+$1A, $00	; #18 (24,40)
	DB $28+$28, $28+$A0, $13+$1A, $00	; #19 (40,40)
	DB $30+$28, $08+$A0, $14+$1A, $00	; #20 ( 8,48)
	DB $30+$28, $10+$A0, $15+$1A, $00	; #21 (16,48)
	DB $30+$28, $18+$A0, $16+$1A, $00	; #22 (24,48)
	DB $30+$28, $20+$A0, $17+$1A, $00	; #23 (32,48)
	DB $30+$28, $28+$A0, $18+$1A, $00	; #24 (40,48)
	DB $38+$28, $18+$A0, $19+$1A, $00	; #25 (24,56)
	DB $38+$28, $20+$A0, $1A+$1A, $00	; #26 (32,56)
	DB $40+$28, $18+$A0, $1B+$1A, $00	; #27 (24,64)
	DB $40+$28, $20+$A0, $1C+$1A, $00	; #28 (32,64)
	DB $48+$28, $18+$A0, $1D+$1A, $00	; #29 (24,72)
	DB $48+$28, $20+$A0, $1E+$1A, $00	; #30 (32,72)
	DB $48+$28, $28+$A0, $1F+$1A, $00	; #31 (40,72)
	DB $50+$28, $20+$A0, $20+$1A, $00	; #32 (32,80)
	DB $50+$28, $28+$A0, $21+$1A, $00	; #33 (40,80)
	DB $50+$28, $30+$A0, $22+$1A, $00	; #34 (48,80)
	DB $58+$28, $20+$A0, $23+$1A, $00	; #35 (32,88)
	DB $58+$28, $28+$A0, $24+$1A, $00	; #36 (40,88)
	DB $60+$28, $28+$A0, $25+$1A, $00	; #37 (40,96)
	DB $68+$28, $28+$A0, $26+$1A, $00	; #38 (40,104)
	DB $70+$28, $28+$A0, $27+$1A, $00	; #39 (40,112)
.end

; Tile data generated with:
; rgbgfx -o tiles.bin ~/Pictures/tilesets/miup/tiles-lean-in.png
tiles_data:
	INCBIN "tiles.bin"
.end