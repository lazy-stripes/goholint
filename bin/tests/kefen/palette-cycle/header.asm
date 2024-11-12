; vim: set fileencoding=utf-8 ts=8 et sw=4 sts=4 :
; Common GameBoy header compatible with rgbds.
; Courtesy of: https://assemblydigest.tumblr.com/post/77198211186/tutorial-making-an-empty-game-boy-rom-in-rgbds

; Useful constants.
INCLUDE "../include/const.asm"

; Interrupt jumps
INCLUDE "../include/interrupts.asm"

SECTION "header", ROM0[$0100]
    ; Main entry point
    NOP
    JP main

    ; $0104 - $0133: The Nintendo Logo.
    DB $CE, $ED, $66, $66, $CC, $0D, $00, $0B
    DB $03, $73, $00, $83, $00, $0C, $00, $0D
    DB $00, $08, $11, $1F, $88, $89, $00, $0E
    DB $DC, $CC, $6E, $E6, $DD, $DD, $D9, $99
    DB $BB, $BB, $67, $63, $6E, $0E, $EC, $CC
    DB $DD, $DC, $99, $9F, $BB, $B9, $33, $3E

    ; $0134 - $013E: The title, in upper-case letters, followed by zeroes.
    DB "PALCYCLE"
    DS 3 ; padding

    ; $013F - $0142: The manufacturer code.
    DS 4

    ; $0143: Gameboy Color compatibility flag.
    DB GBC_UNSUPPORTED

    ; $0144 - $0145: "New" Licensee Code, a two character name.
    DB "KF"

    ; $0146: Super Gameboy compatibility flag.
    DB SGB_UNSUPPORTED

    ; $0147: Cartridge type. Either no ROM or MBC5 is recommended.
    DB CART_ROM_ONLY

    ; $0148: Rom size.
    DB ROM_32K

    ; $0149: Ram size.
    DB RAM_NONE

    ; $014A: Destination code.
    DB DEST_INTERNATIONAL

    ; $014B: Old licensee code.
    ; $33 indicates new license code will be used.
    ; $33 must be used for SGB games.
    DB $33

    ; $014C: ROM version number
    DB $00

    ; $014D: Header checksum.
    ; Assembler needs to patch this.
    DB $00

    ; $014E- $014F: Global checksum.
    ; Assembler needs to patch this.
    DW $0000

    ; Done. Write actual code with a main: label in another file.
