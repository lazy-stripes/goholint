; Standard interrupt space just jumping to the relevant label.
SECTION "interrupts", ROM0[$0040]
    ; $0040 - $0067: Interrupt handlers.
    JP vblank
    REPT 5
        NOP
    ENDR
    ; $0048
    JP stat
    REPT 5
        NOP 
    ENDR
    ; $0050
    JP timer
    REPT 5
        NOP
    ENDR
    ; $0058
    JP serial
    REPT 5
        NOP
    ENDR
    ; $0060
    JP joypad
    REPT 5
        NOP
    ENDR