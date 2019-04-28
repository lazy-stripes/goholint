; Tiny test to try rgbds. Break scroll and see...

INCLUDE "header.asm"

SECTION "default", ROM0
main:
init_oam:
	; Initialize shadow OAM
	LD DE, sprites_data
	LD HL, $C000
	LD C, sprites_data.end - sprites_data
.copy_loop
	LD A, [DE]
	INC DE
	LD [HLI], A
	DEC C
	JR NZ, .copy_loop

init_tiles:
	; XXX: turning LCD off for now.
	;XOR A
	;LD [$FF00+$40],A

	; Copy tile data to VRAM.
	LD DE, tiles_data
	LD HL, $81A0	; Right after Nintendo logo tiles
	LD BC, tiles_data.end - tiles_data
.copy_loop
	LD A, [DE]
	INC DE
	LD [HLI], A
	DEC BC	; Does not set zero flag so we need to check
	LD A, B
	CP 0
	JR NZ, .copy_loop
	LD A, C
	CP 0
	JR NZ, .copy_loop

	;LD A, $91
	;LD [$FF00+$40],A

init_dma:
	CALL copy_dma_routine

	; Init state and enable vblank interrupt.
	LD HL, state_update_sprites
	LD A, $01
	LD [$FF00+$FF], A
	EI

game_loop:
	JR game_loop


;
; Interrupts
;

; Call game state defined in HL
vblank:
	; Emulating a CALL HL of sorts. I'm probably doing this wrong.
	LD BC, .end
	PUSH BC	; Store return address in stack
	JP HL
.end
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
state_update_sprites:
	; TODO: actually update OAM shadow RAM

	CALL $FF80	; DMA
	; TODO: next state: LD HL, state_idle_wait
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
copy_dma_routine:
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
	; Using $1A offset to put our tiles right after the Nintendo logo.
	;   Y    X    #    Palette
	DB $10, $08, $00+$1A, $00	; #0  ( 8,16)
	DB $10, $10, $01+$1A, $00	; #1  (16,16)
	DB $10, $18, $02+$1A, $00	; #2  (24,16)
	DB $10, $20, $03+$1A, $00	; #3  (32,16)
	DB $10, $28, $04+$1A, $00	; #4  (40,16)
	DB $18, $08, $05+$1A, $00	; #5  ( 8,24)
	DB $18, $10, $06+$1A, $00	; #6  (16,24)
	DB $18, $18, $07+$1A, $00	; #7  (24,24)
	DB $18, $20, $08+$1A, $00	; #8  (32,24)
	DB $18, $28, $09+$1A, $00	; #9  (40,24)
	DB $20, $08, $0A+$1A, $00	; #10 ( 8,32)
	DB $20, $10, $0B+$1A, $00	; #11 (16,32)
	DB $20, $18, $0C+$1A, $00	; #12 (24,32)
	DB $20, $20, $0D+$1A, $00	; #13 (32,32)
	DB $20, $28, $0E+$1A, $00	; #14 (40,32)
	DB $20, $30, $0F+$1A, $00	; #15 (48,32)
	DB $28, $08, $10+$1A, $00	; #16 ( 8,40)
	DB $28, $10, $11+$1A, $00	; #17 (16,40)
	DB $28, $18, $12+$1A, $00	; #18 (24,40)
	DB $28, $28, $13+$1A, $00	; #19 (40,40)
	DB $30, $08, $14+$1A, $00	; #20 ( 8,48)
	DB $30, $10, $15+$1A, $00	; #21 (16,48)
	DB $30, $18, $16+$1A, $00	; #22 (24,48)
	DB $30, $20, $17+$1A, $00	; #23 (32,48)
	DB $30, $28, $18+$1A, $00	; #24 (40,48)
	DB $38, $18, $19+$1A, $00	; #25 (24,56)
	DB $38, $20, $1A+$1A, $00	; #26 (32,56)
	DB $40, $18, $1B+$1A, $00	; #27 (24,64)
	DB $40, $20, $1C+$1A, $00	; #28 (32,64)
	DB $48, $18, $1D+$1A, $00	; #29 (24,72)
	DB $48, $20, $1E+$1A, $00	; #30 (32,72)
	DB $48, $28, $1F+$1A, $00	; #31 (40,72)
	DB $50, $20, $20+$1A, $00	; #32 (32,80)
	DB $50, $28, $21+$1A, $00	; #33 (40,80)
	DB $50, $30, $22+$1A, $00	; #34 (48,80)
	DB $58, $20, $23+$1A, $00	; #35 (32,88)
	DB $58, $28, $24+$1A, $00	; #36 (40,88)
	DB $60, $28, $25+$1A, $00	; #37 (40,96)
	DB $68, $28, $26+$1A, $00	; #38 (40,104)
	DB $70, $28, $27+$1A, $00	; #39 (40,112)
.end

; Tile data generated with:
; rgbgfx -o tiles.bin ~/Pictures/tilesets/miup/tiles-lean-in.png
tiles_data:
	INCBIN "tiles.bin"
.end