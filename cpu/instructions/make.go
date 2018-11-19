package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"
)

// Quick and dirty aggregate of keys, not necessarily all used by all instructions, for readability.
type data struct {
	Extended      bool
	Opcode        uint
	Bit           uint
	Template      string
	Instruction   string
	High, Low     string
	Flag          string
	Register      string
	OtherRegister string
	Address       string
	Operator      string
}

type instructionSet struct {
	Instructions []data
	Extended     []data
}

// typeName generates a deterministic name for derived instruction types.
// TODO: I'm going the easy way for now, but I'd love being able to generate
//		 human-readable name like nop, ldDeD16, etc.
func typeName(d *data) (name string) {
	var prefix string
	if d.Extended {
		prefix = "opCb"
	} else {
		prefix = "op"
	}
	return fmt.Sprintf("%s%s", prefix, strings.Title(fmt.Sprintf("%02x", d.Opcode)))
}

func main() {
	instructions := []data{
		{Opcode: 0x00, Template: "nop"},
		{Opcode: 0x01, Template: "ldrrd16", High: "B", Low: "C"},
		{Opcode: 0x02, Template: "ldaddrr", Address: "BC", Register: "A"},
		{Opcode: 0x03, Template: "incrr", High: "B", Low: "C"},
		{Opcode: 0x04, Template: "incr", Register: "B"},
		{Opcode: 0x05, Template: "decr", Register: "B"},
		{Opcode: 0x06, Template: "ldrd8", Register: "B"},
		{Opcode: 0x07, Template: "rlca", Register: "A"},
		{Opcode: 0x08, Template: "lda16sp"},
		{Opcode: 0x09, Template: "addhlrr", High: "B", Low: "C"},
		{Opcode: 0x0a, Template: "ldaddrr", Register: "A", Address: "BC"},
		{Opcode: 0x0b, Template: "decrr", High: "B", Low: "C"},
		{Opcode: 0x0c, Template: "incr", Register: "C"},
		{Opcode: 0x0d, Template: "decr", Register: "C"},
		{Opcode: 0x0e, Template: "ldrd8", Register: "C"},

		{Opcode: 0x10, Template: "stop"},
		{Opcode: 0x11, Template: "ldrrd16", High: "D", Low: "E"},
		{Opcode: 0x12, Template: "ldaddrr", Address: "DE", Register: "A"},
		{Opcode: 0x13, Template: "incrr", High: "D", Low: "E"},
		{Opcode: 0x14, Template: "incr", Register: "D"},
		{Opcode: 0x15, Template: "decr", Register: "D"},
		{Opcode: 0x16, Template: "ldrd8", Register: "D"},
		{Opcode: 0x17, Template: "rla", Register: "A"},
		{Opcode: 0x18, Template: "jr"},
		{Opcode: 0x19, Template: "addhlrr", High: "D", Low: "E"},
		{Opcode: 0x1a, Template: "ldraddr", Register: "A", Address: "DE"},
		{Opcode: 0x1b, Template: "decrr", High: "D", Low: "E"},
		{Opcode: 0x1c, Template: "incr", Register: "E"},
		{Opcode: 0x1d, Template: "decr", Register: "E"},
		{Opcode: 0x1e, Template: "ldrd8", Register: "E"},
		{Opcode: 0x1f, Template: "rra", Register: "A"},
		{Opcode: 0x20, Template: "jr", Operator: "!", Flag: "Z"},
		{Opcode: 0x21, Template: "ldrrd16", High: "H", Low: "L"},
		{Opcode: 0x22, Template: "ldidhla", Operator: "+"},
		{Opcode: 0x23, Template: "incrr", High: "H", Low: "L"},
		{Opcode: 0x24, Template: "incr", Register: "H"},
		{Opcode: 0x25, Template: "decr", Register: "H"},
		{Opcode: 0x26, Template: "ldrd8", Register: "H"},

		{Opcode: 0x28, Template: "jr", Operator: "=", Flag: "Z"},
		{Opcode: 0x29, Template: "addhlrr", High: "H", Low: "L"},
		{Opcode: 0x2a, Template: "ldidahl", Operator: "+"},
		{Opcode: 0x2b, Template: "decrr", High: "H", Low: "L"},
		{Opcode: 0x2c, Template: "incr", Register: "L"},
		{Opcode: 0x2d, Template: "decr", Register: "L"},
		{Opcode: 0x2e, Template: "ldrd8", Register: "L"},
		{Opcode: 0x2f, Template: "cpl"},
		{Opcode: 0x30, Template: "jr", Operator: "!", Flag: "C"},
		{Opcode: 0x31, Template: "ldrrd16", High: "S", Low: "P"},
		{Opcode: 0x32, Template: "ldidhla", Operator: "-"},
		{Opcode: 0x33, Template: "incrr", High: "S", Low: "P"},
		{Opcode: 0x34, Template: "incaddr", Address: "HL"},
		{Opcode: 0x35, Template: "decaddr", Address: "HL"},
		{Opcode: 0x36, Template: "ldaddrd8", Address: "HL"},

		{Opcode: 0x38, Template: "jr", Operator: "=", Flag: "C"},
		{Opcode: 0x39, Template: "addhlrr", High: "S", Low: "P"},
		{Opcode: 0x3a, Template: "ldidahl", Operator: "-"},
		{Opcode: 0x3b, Template: "decrr", High: "S", Low: "P"},
		{Opcode: 0x3c, Template: "incr", Register: "A"},
		{Opcode: 0x3d, Template: "decr", Register: "A"},
		{Opcode: 0x3e, Template: "ldrd8", Register: "A"},

		{Opcode: 0x40, Template: "ldrr", Register: "B", OtherRegister: "B"},
		{Opcode: 0x41, Template: "ldrr", Register: "B", OtherRegister: "C"},
		{Opcode: 0x42, Template: "ldrr", Register: "B", OtherRegister: "D"},
		{Opcode: 0x43, Template: "ldrr", Register: "B", OtherRegister: "E"},
		{Opcode: 0x44, Template: "ldrr", Register: "B", OtherRegister: "H"},
		{Opcode: 0x45, Template: "ldrr", Register: "B", OtherRegister: "L"},
		{Opcode: 0x46, Template: "ldraddr", Register: "B", Address: "HL"},
		{Opcode: 0x47, Template: "ldrr", Register: "B", OtherRegister: "A"},
		{Opcode: 0x48, Template: "ldrr", Register: "C", OtherRegister: "B"},
		{Opcode: 0x49, Template: "ldrr", Register: "C", OtherRegister: "C"},
		{Opcode: 0x4a, Template: "ldrr", Register: "C", OtherRegister: "D"},
		{Opcode: 0x4b, Template: "ldrr", Register: "C", OtherRegister: "E"},
		{Opcode: 0x4c, Template: "ldrr", Register: "C", OtherRegister: "H"},
		{Opcode: 0x4d, Template: "ldrr", Register: "C", OtherRegister: "L"},
		{Opcode: 0x4e, Template: "ldraddr", Register: "C", Address: "HL"},
		{Opcode: 0x4f, Template: "ldrr", Register: "C", OtherRegister: "A"},
		{Opcode: 0x50, Template: "ldrr", Register: "D", OtherRegister: "B"},
		{Opcode: 0x51, Template: "ldrr", Register: "D", OtherRegister: "C"},
		{Opcode: 0x52, Template: "ldrr", Register: "D", OtherRegister: "D"},
		{Opcode: 0x53, Template: "ldrr", Register: "D", OtherRegister: "E"},
		{Opcode: 0x54, Template: "ldrr", Register: "D", OtherRegister: "H"},
		{Opcode: 0x55, Template: "ldrr", Register: "D", OtherRegister: "L"},
		{Opcode: 0x56, Template: "ldraddr", Register: "D", Address: "HL"},
		{Opcode: 0x57, Template: "ldrr", Register: "D", OtherRegister: "A"},
		{Opcode: 0x58, Template: "ldrr", Register: "E", OtherRegister: "B"},
		{Opcode: 0x59, Template: "ldrr", Register: "E", OtherRegister: "C"},
		{Opcode: 0x5a, Template: "ldrr", Register: "E", OtherRegister: "D"},
		{Opcode: 0x5b, Template: "ldrr", Register: "E", OtherRegister: "E"},
		{Opcode: 0x5c, Template: "ldrr", Register: "E", OtherRegister: "H"},
		{Opcode: 0x5d, Template: "ldrr", Register: "E", OtherRegister: "L"},
		{Opcode: 0x5e, Template: "ldraddr", Register: "E", Address: "HL"},
		{Opcode: 0x5f, Template: "ldrr", Register: "E", OtherRegister: "A"},
		{Opcode: 0x60, Template: "ldrr", Register: "H", OtherRegister: "B"},
		{Opcode: 0x61, Template: "ldrr", Register: "H", OtherRegister: "C"},
		{Opcode: 0x62, Template: "ldrr", Register: "H", OtherRegister: "D"},
		{Opcode: 0x63, Template: "ldrr", Register: "H", OtherRegister: "E"},
		{Opcode: 0x64, Template: "ldrr", Register: "H", OtherRegister: "H"},
		{Opcode: 0x65, Template: "ldrr", Register: "H", OtherRegister: "L"},
		{Opcode: 0x66, Template: "ldraddr", Register: "H", Address: "HL"},
		{Opcode: 0x67, Template: "ldrr", Register: "H", OtherRegister: "A"},
		{Opcode: 0x68, Template: "ldrr", Register: "L", OtherRegister: "B"},
		{Opcode: 0x69, Template: "ldrr", Register: "L", OtherRegister: "C"},
		{Opcode: 0x6a, Template: "ldrr", Register: "L", OtherRegister: "D"},
		{Opcode: 0x6b, Template: "ldrr", Register: "L", OtherRegister: "E"},
		{Opcode: 0x6c, Template: "ldrr", Register: "L", OtherRegister: "H"},
		{Opcode: 0x6d, Template: "ldrr", Register: "L", OtherRegister: "L"},
		{Opcode: 0x6e, Template: "ldraddr", Register: "L", Address: "HL"},
		{Opcode: 0x6f, Template: "ldrr", Register: "L", OtherRegister: "A"},
		{Opcode: 0x70, Template: "ldaddrr", Address: "HL", Register: "B"},
		{Opcode: 0x71, Template: "ldaddrr", Address: "HL", Register: "C"},
		{Opcode: 0x72, Template: "ldaddrr", Address: "HL", Register: "D"},
		{Opcode: 0x73, Template: "ldaddrr", Address: "HL", Register: "E"},
		{Opcode: 0x74, Template: "ldaddrr", Address: "HL", Register: "H"},
		{Opcode: 0x75, Template: "ldaddrr", Address: "HL", Register: "L"},
		{Opcode: 0x76, Template: "halt"},
		{Opcode: 0x77, Template: "ldaddrr", Address: "HL", Register: "A"},
		{Opcode: 0x78, Template: "ldrr", Register: "A", OtherRegister: "B"},
		{Opcode: 0x79, Template: "ldrr", Register: "A", OtherRegister: "C"},
		{Opcode: 0x7a, Template: "ldrr", Register: "A", OtherRegister: "D"},
		{Opcode: 0x7b, Template: "ldrr", Register: "A", OtherRegister: "E"},
		{Opcode: 0x7c, Template: "ldrr", Register: "A", OtherRegister: "H"},
		{Opcode: 0x7d, Template: "ldrr", Register: "A", OtherRegister: "L"},
		{Opcode: 0x7e, Template: "ldraddr", Register: "A", Address: "HL"},
		{Opcode: 0x7f, Template: "ldrr", Register: "A", OtherRegister: "A"},
		{Opcode: 0x80, Template: "addr", Register: "B"},
		{Opcode: 0x81, Template: "addr", Register: "C"},
		{Opcode: 0x82, Template: "addr", Register: "D"},
		{Opcode: 0x83, Template: "addr", Register: "E"},
		{Opcode: 0x84, Template: "addr", Register: "H"},
		{Opcode: 0x85, Template: "addr", Register: "L"},
		{Opcode: 0x86, Template: "addaddr", Address: "HL"},
		{Opcode: 0x87, Template: "addr", Register: "A"},

		{Opcode: 0x90, Template: "subcpr", Instruction: "SUB", Register: "B"},
		{Opcode: 0x91, Template: "subcpr", Instruction: "SUB", Register: "C"},
		{Opcode: 0x92, Template: "subcpr", Instruction: "SUB", Register: "D"},
		{Opcode: 0x93, Template: "subcpr", Instruction: "SUB", Register: "E"},
		{Opcode: 0x94, Template: "subcpr", Instruction: "SUB", Register: "H"},
		{Opcode: 0x95, Template: "subcpr", Instruction: "SUB", Register: "L"},
		{Opcode: 0x96, Template: "subcpaddr", Instruction: "SUB", Address: "HL"},
		{Opcode: 0x97, Template: "subcpr", Instruction: "SUB", Register: "A"},
		{Opcode: 0x98, Template: "sbcr", Register: "B"},
		{Opcode: 0x99, Template: "sbcr", Register: "C"},
		{Opcode: 0x9a, Template: "sbcr", Register: "D"},
		{Opcode: 0x9b, Template: "sbcr", Register: "E"},
		{Opcode: 0x9c, Template: "sbcr", Register: "H"},
		{Opcode: 0x9d, Template: "sbcr", Register: "L"},
		//{Opcode: 0x9e, Template: "sbcaddr", Address: "HL"},
		{Opcode: 0x9f, Template: "sbcr", Register: "A"},

		{Opcode: 0xa0, Template: "boolr", Instruction: "AND", Operator: "&=", Register: "B"},
		{Opcode: 0xa1, Template: "boolr", Instruction: "AND", Operator: "&=", Register: "C"},
		{Opcode: 0xa2, Template: "boolr", Instruction: "AND", Operator: "&=", Register: "D"},
		{Opcode: 0xa3, Template: "boolr", Instruction: "AND", Operator: "&=", Register: "E"},
		{Opcode: 0xa4, Template: "boolr", Instruction: "AND", Operator: "&=", Register: "H"},
		{Opcode: 0xa5, Template: "boolr", Instruction: "AND", Operator: "&=", Register: "L"},
		{Opcode: 0xa6, Template: "booladdr", Instruction: "AND", Operator: "&=", Address: "HL"},
		{Opcode: 0xa7, Template: "boolr", Instruction: "AND", Operator: "&=", Register: "A"},
		{Opcode: 0xa8, Template: "boolr", Instruction: "XOR", Operator: "^=", Register: "B"},
		{Opcode: 0xa9, Template: "boolr", Instruction: "XOR", Operator: "^=", Register: "C"},
		{Opcode: 0xaa, Template: "boolr", Instruction: "XOR", Operator: "^=", Register: "D"},
		{Opcode: 0xab, Template: "boolr", Instruction: "XOR", Operator: "^=", Register: "E"},
		{Opcode: 0xac, Template: "boolr", Instruction: "XOR", Operator: "^=", Register: "H"},
		{Opcode: 0xad, Template: "boolr", Instruction: "XOR", Operator: "^=", Register: "L"},
		{Opcode: 0xae, Template: "booladdr", Instruction: "XOR", Operator: "^=", Address: "HL"},
		{Opcode: 0xaf, Template: "boolr", Instruction: "XOR", Operator: "^=", Register: "A"},
		{Opcode: 0xb0, Template: "boolr", Instruction: "OR", Operator: "|=", Register: "B"},
		{Opcode: 0xb1, Template: "boolr", Instruction: "OR", Operator: "|=", Register: "C"},
		{Opcode: 0xb2, Template: "boolr", Instruction: "OR", Operator: "|=", Register: "D"},
		{Opcode: 0xb3, Template: "boolr", Instruction: "OR", Operator: "|=", Register: "E"},
		{Opcode: 0xb4, Template: "boolr", Instruction: "OR", Operator: "|=", Register: "H"},
		{Opcode: 0xb5, Template: "boolr", Instruction: "OR", Operator: "|=", Register: "L"},
		{Opcode: 0xb6, Template: "booladdr", Instruction: "OR", Operator: "|=", Address: "HL"},
		{Opcode: 0xb7, Template: "boolr", Instruction: "OR", Operator: "|=", Register: "A"},
		{Opcode: 0xb8, Template: "subcpr", Instruction: "CP", Register: "B"},
		{Opcode: 0xb9, Template: "subcpr", Instruction: "CP", Register: "C"},
		{Opcode: 0xba, Template: "subcpr", Instruction: "CP", Register: "D"},
		{Opcode: 0xbb, Template: "subcpr", Instruction: "CP", Register: "E"},
		{Opcode: 0xbc, Template: "subcpr", Instruction: "CP", Register: "H"},
		{Opcode: 0xbd, Template: "subcpr", Instruction: "CP", Register: "L"},
		{Opcode: 0xbe, Template: "subcpaddr", Instruction: "CP", Address: "HL"},
		{Opcode: 0xbf, Template: "subcpr", Instruction: "CP", Register: "A"},
		{Opcode: 0xc0, Template: "ret", Instruction: "RET", Operator: "!", Flag: "Z"},
		{Opcode: 0xc1, Template: "pop", High: "B", Low: "C"},

		{Opcode: 0xc2, Template: "calljp", Instruction: "JP", Operator: "!", Flag: "Z"},
		{Opcode: 0xc3, Template: "calljp", Instruction: "JP"},
		{Opcode: 0xc4, Template: "calljp", Instruction: "CALL", Operator: "!", Flag: "Z"},
		{Opcode: 0xc5, Template: "push", High: "B", Low: "C"},
		{Opcode: 0xc6, Template: "addaddr"},
		{Opcode: 0xc7, Template: "rst", Address: "00"},

		{Opcode: 0xc8, Template: "ret", Instruction: "RET", Operator: "=", Flag: "Z"},
		{Opcode: 0xc9, Template: "ret", Instruction: "RET"},
		{Opcode: 0xca, Template: "calljp", Instruction: "JP", Operator: "=", Flag: "Z"},
		// CB is prefix for extended instruction set.
		{Opcode: 0xcc, Template: "calljp", Instruction: "CALL", Operator: "=", Flag: "Z"},
		{Opcode: 0xcd, Template: "calljp", Instruction: "CALL"},
		{Opcode: 0xce, Template: "adcaddr"},
		{Opcode: 0xcf, Template: "rst", Address: "08"},
		{Opcode: 0xd0, Template: "ret", Instruction: "RET", Operator: "!", Flag: "C"},
		{Opcode: 0xd1, Template: "pop", High: "D", Low: "E"},
		{Opcode: 0xd2, Template: "calljp", Instruction: "JP", Operator: "!", Flag: "C"},
		// No D3 opcode.
		{Opcode: 0xd4, Template: "calljp", Instruction: "CALL", Operator: "!", Flag: "C"},
		{Opcode: 0xd5, Template: "push", High: "D", Low: "E"},
		{Opcode: 0xd6, Template: "subcpaddr", Instruction: "SUB"},
		{Opcode: 0xd7, Template: "rst", Address: "10"},
		{Opcode: 0xd8, Template: "ret", Instruction: "RET", Operator: "=", Flag: "C"},

		{Opcode: 0xd9, Template: "ret", Instruction: "RETI"},
		{Opcode: 0xda, Template: "calljp", Instruction: "JP", Operator: "=", Flag: "C"},
		// No DB opcode.
		{Opcode: 0xdc, Template: "calljp", Instruction: "CALL", Operator: "=", Flag: "C"},
		{Opcode: 0xde, Template: "sbcaddr"},
		{Opcode: 0xdf, Template: "rst", Address: "18"},
		{Opcode: 0xe0, Template: "ldioa"},
		{Opcode: 0xe1, Template: "pop", High: "H", Low: "L"},
		{Opcode: 0xe2, Template: "ldioa", Register: "C"},

		{Opcode: 0xe5, Template: "push", High: "H", Low: "L"},
		{Opcode: 0xe6, Template: "booladdr", Instruction: "AND", Operator: "&="},
		{Opcode: 0xe7, Template: "rst", Address: "20"},
		{Opcode: 0xe8, Template: "addspr8", Register: "SP"},
		{Opcode: 0xe9, Template: "jphl"},

		{Opcode: 0xea, Template: "lda16a"},
		{Opcode: 0xee, Template: "booladdr", Instruction: "XOR", Operator: "^="},
		{Opcode: 0xef, Template: "rst", Address: "28"},

		{Opcode: 0xf0, Template: "ldaio"},
		{Opcode: 0xf1, Template: "pop", High: "A", Low: "F"},

		{Opcode: 0xf3, Template: "interrupt", Instruction: "DI"},

		{Opcode: 0xf5, Template: "push", High: "A", Low: "F"},
		{Opcode: 0xf6, Template: "booladdr", Instruction: "OR", Operator: "|="},
		{Opcode: 0xf7, Template: "rst", Address: "30"},
		{Opcode: 0xf8, Template: "addspr8", Register: "HL"},
		{Opcode: 0xf9, Template: "ldsphl"},

		{Opcode: 0xfa, Template: "ldaa16"},
		{Opcode: 0xfb, Template: "interrupt", Instruction: "EI"},

		{Opcode: 0xfe, Template: "subcpaddr", Instruction: "CP"},
		{Opcode: 0xff, Template: "rst", Address: "38"},
	}

	extended := []data{
		// FIXME: those should be dynamically generated there and then.
		{Extended: true, Opcode: 0x10, Template: "rlr", Register: "B"},
		{Extended: true, Opcode: 0x11, Template: "rlr", Register: "C"},
		{Extended: true, Opcode: 0x12, Template: "rlr", Register: "D"},
		{Extended: true, Opcode: 0x13, Template: "rlr", Register: "E"},
		{Extended: true, Opcode: 0x14, Template: "rlr", Register: "H"},
		{Extended: true, Opcode: 0x15, Template: "rlr", Register: "L"},
		//{Extended: true, Opcode: 0x16, Template: "rladdr", Address: "HL"},
		{Extended: true, Opcode: 0x17, Template: "rlr", Register: "A"},
		{Extended: true, Opcode: 0x18, Template: "rrr", Register: "B"},
		{Extended: true, Opcode: 0x19, Template: "rrr", Register: "C"},
		{Extended: true, Opcode: 0x1a, Template: "rrr", Register: "D"},
		{Extended: true, Opcode: 0x1b, Template: "rrr", Register: "E"},
		{Extended: true, Opcode: 0x1c, Template: "rrr", Register: "H"},
		{Extended: true, Opcode: 0x1d, Template: "rrr", Register: "L"},
		//{Extended: true, Opcode: 0x1e, Template: "rraddr", Register: "HL"},
		{Extended: true, Opcode: 0x1f, Template: "rrr", Register: "A"},

		{Extended: true, Opcode: 0x30, Template: "swapr", Register: "B"},
		{Extended: true, Opcode: 0x31, Template: "swapr", Register: "C"},
		{Extended: true, Opcode: 0x32, Template: "swapr", Register: "D"},
		{Extended: true, Opcode: 0x33, Template: "swapr", Register: "E"},
		{Extended: true, Opcode: 0x34, Template: "swapr", Register: "H"},
		{Extended: true, Opcode: 0x35, Template: "swapr", Register: "L"},
		//{Extended: true, Opcode: 0x36, Template: "swapaddr", Register: "HL"},
		{Extended: true, Opcode: 0x37, Template: "swapr", Register: "A"},
		{Extended: true, Opcode: 0x38, Template: "sr", Instruction: "SRL", Register: "B"},
		{Extended: true, Opcode: 0x39, Template: "sr", Instruction: "SRL", Register: "C"},
		{Extended: true, Opcode: 0x3a, Template: "sr", Instruction: "SRL", Register: "D"},
		{Extended: true, Opcode: 0x3b, Template: "sr", Instruction: "SRL", Register: "E"},
		{Extended: true, Opcode: 0x3c, Template: "sr", Instruction: "SRL", Register: "H"},
		{Extended: true, Opcode: 0x3d, Template: "sr", Instruction: "SRL", Register: "L"},
		//{Extended: true, Opcode: 0x3e, Template: "sr", Instruction: "SRL", Register: "HL"},
		{Extended: true, Opcode: 0x3f, Template: "sr", Instruction: "SRL", Register: "A"},
		{Extended: true, Opcode: 0x40, Template: "bitnr", Bit: 0, Register: "B"},
		{Extended: true, Opcode: 0x41, Template: "bitnr", Bit: 0, Register: "C"},
		{Extended: true, Opcode: 0x42, Template: "bitnr", Bit: 0, Register: "D"},
		{Extended: true, Opcode: 0x43, Template: "bitnr", Bit: 0, Register: "E"},
		{Extended: true, Opcode: 0x44, Template: "bitnr", Bit: 0, Register: "H"},
		{Extended: true, Opcode: 0x45, Template: "bitnr", Bit: 0, Register: "L"},
		{Extended: true, Opcode: 0x46, Template: "bitnaddr", Bit: 0, Address: "HL"},
		{Extended: true, Opcode: 0x47, Template: "bitnr", Bit: 0, Register: "A"},
		{Extended: true, Opcode: 0x48, Template: "bitnr", Bit: 1, Register: "B"},
		{Extended: true, Opcode: 0x49, Template: "bitnr", Bit: 1, Register: "C"},
		{Extended: true, Opcode: 0x4a, Template: "bitnr", Bit: 1, Register: "D"},
		{Extended: true, Opcode: 0x4b, Template: "bitnr", Bit: 1, Register: "E"},
		{Extended: true, Opcode: 0x4c, Template: "bitnr", Bit: 1, Register: "H"},
		{Extended: true, Opcode: 0x4d, Template: "bitnr", Bit: 1, Register: "L"},
		{Extended: true, Opcode: 0x4e, Template: "bitnaddr", Bit: 1, Address: "HL"},
		{Extended: true, Opcode: 0x4f, Template: "bitnr", Bit: 1, Register: "A"},
		{Extended: true, Opcode: 0x50, Template: "bitnr", Bit: 2, Register: "B"},
		{Extended: true, Opcode: 0x51, Template: "bitnr", Bit: 2, Register: "C"},
		{Extended: true, Opcode: 0x52, Template: "bitnr", Bit: 2, Register: "D"},
		{Extended: true, Opcode: 0x53, Template: "bitnr", Bit: 2, Register: "E"},
		{Extended: true, Opcode: 0x54, Template: "bitnr", Bit: 2, Register: "H"},
		{Extended: true, Opcode: 0x55, Template: "bitnr", Bit: 2, Register: "L"},
		{Extended: true, Opcode: 0x56, Template: "bitnaddr", Bit: 2, Address: "HL"},
		{Extended: true, Opcode: 0x57, Template: "bitnr", Bit: 2, Register: "A"},
		{Extended: true, Opcode: 0x58, Template: "bitnr", Bit: 3, Register: "B"},
		{Extended: true, Opcode: 0x59, Template: "bitnr", Bit: 3, Register: "C"},
		{Extended: true, Opcode: 0x5a, Template: "bitnr", Bit: 3, Register: "D"},
		{Extended: true, Opcode: 0x5b, Template: "bitnr", Bit: 3, Register: "E"},
		{Extended: true, Opcode: 0x5c, Template: "bitnr", Bit: 3, Register: "H"},
		{Extended: true, Opcode: 0x5d, Template: "bitnr", Bit: 3, Register: "L"},
		{Extended: true, Opcode: 0x5e, Template: "bitnaddr", Bit: 3, Address: "HL"},
		{Extended: true, Opcode: 0x5f, Template: "bitnr", Bit: 3, Register: "A"},
		{Extended: true, Opcode: 0x60, Template: "bitnr", Bit: 4, Register: "B"},
		{Extended: true, Opcode: 0x61, Template: "bitnr", Bit: 4, Register: "C"},
		{Extended: true, Opcode: 0x62, Template: "bitnr", Bit: 4, Register: "D"},
		{Extended: true, Opcode: 0x63, Template: "bitnr", Bit: 4, Register: "E"},
		{Extended: true, Opcode: 0x64, Template: "bitnr", Bit: 4, Register: "H"},
		{Extended: true, Opcode: 0x65, Template: "bitnr", Bit: 4, Register: "L"},
		{Extended: true, Opcode: 0x66, Template: "bitnaddr", Bit: 4, Address: "HL"},
		{Extended: true, Opcode: 0x67, Template: "bitnr", Bit: 4, Register: "A"},
		{Extended: true, Opcode: 0x68, Template: "bitnr", Bit: 5, Register: "B"},
		{Extended: true, Opcode: 0x69, Template: "bitnr", Bit: 5, Register: "C"},
		{Extended: true, Opcode: 0x6a, Template: "bitnr", Bit: 5, Register: "D"},
		{Extended: true, Opcode: 0x6b, Template: "bitnr", Bit: 5, Register: "E"},
		{Extended: true, Opcode: 0x6c, Template: "bitnr", Bit: 5, Register: "H"},
		{Extended: true, Opcode: 0x6d, Template: "bitnr", Bit: 5, Register: "L"},
		{Extended: true, Opcode: 0x6e, Template: "bitnaddr", Bit: 5, Address: "HL"},
		{Extended: true, Opcode: 0x6f, Template: "bitnr", Bit: 5, Register: "A"},
		{Extended: true, Opcode: 0x70, Template: "bitnr", Bit: 6, Register: "B"},
		{Extended: true, Opcode: 0x71, Template: "bitnr", Bit: 6, Register: "C"},
		{Extended: true, Opcode: 0x72, Template: "bitnr", Bit: 6, Register: "D"},
		{Extended: true, Opcode: 0x73, Template: "bitnr", Bit: 6, Register: "E"},
		{Extended: true, Opcode: 0x74, Template: "bitnr", Bit: 6, Register: "H"},
		{Extended: true, Opcode: 0x75, Template: "bitnr", Bit: 6, Register: "L"},
		{Extended: true, Opcode: 0x76, Template: "bitnaddr", Bit: 6, Address: "HL"},
		{Extended: true, Opcode: 0x77, Template: "bitnr", Bit: 6, Register: "A"},
		{Extended: true, Opcode: 0x78, Template: "bitnr", Bit: 7, Register: "B"},
		{Extended: true, Opcode: 0x79, Template: "bitnr", Bit: 7, Register: "C"},
		{Extended: true, Opcode: 0x7a, Template: "bitnr", Bit: 7, Register: "D"},
		{Extended: true, Opcode: 0x7b, Template: "bitnr", Bit: 7, Register: "E"},
		{Extended: true, Opcode: 0x7c, Template: "bitnr", Bit: 7, Register: "H"},
		{Extended: true, Opcode: 0x7d, Template: "bitnr", Bit: 7, Register: "L"},
		{Extended: true, Opcode: 0x7e, Template: "bitnaddr", Bit: 7, Address: "HL"},
		{Extended: true, Opcode: 0x7f, Template: "bitnr", Bit: 7, Register: "A"},
		{Extended: true, Opcode: 0x80, Template: "resnr", Bit: 0, Register: "B"},
		{Extended: true, Opcode: 0x81, Template: "resnr", Bit: 0, Register: "C"},
		{Extended: true, Opcode: 0x82, Template: "resnr", Bit: 0, Register: "D"},
		{Extended: true, Opcode: 0x83, Template: "resnr", Bit: 0, Register: "E"},
		{Extended: true, Opcode: 0x84, Template: "resnr", Bit: 0, Register: "H"},
		{Extended: true, Opcode: 0x85, Template: "resnr", Bit: 0, Register: "L"},
		{Extended: true, Opcode: 0x86, Template: "resnaddr", Bit: 0, Register: "HL"},
		{Extended: true, Opcode: 0x87, Template: "resnr", Bit: 0, Register: "A"},
		{Extended: true, Opcode: 0x88, Template: "resnr", Bit: 1, Register: "B"},
		{Extended: true, Opcode: 0x89, Template: "resnr", Bit: 1, Register: "C"},
		{Extended: true, Opcode: 0x8a, Template: "resnr", Bit: 1, Register: "D"},
		{Extended: true, Opcode: 0x8b, Template: "resnr", Bit: 1, Register: "E"},
		{Extended: true, Opcode: 0x8c, Template: "resnr", Bit: 1, Register: "H"},
		{Extended: true, Opcode: 0x8d, Template: "resnr", Bit: 1, Register: "L"},
		{Extended: true, Opcode: 0x8e, Template: "resnaddr", Bit: 1, Register: "HL"},
		{Extended: true, Opcode: 0x8f, Template: "resnr", Bit: 1, Register: "A"},
		{Extended: true, Opcode: 0x90, Template: "resnr", Bit: 2, Register: "B"},
		{Extended: true, Opcode: 0x91, Template: "resnr", Bit: 2, Register: "C"},
		{Extended: true, Opcode: 0x92, Template: "resnr", Bit: 2, Register: "D"},
		{Extended: true, Opcode: 0x93, Template: "resnr", Bit: 2, Register: "E"},
		{Extended: true, Opcode: 0x94, Template: "resnr", Bit: 2, Register: "H"},
		{Extended: true, Opcode: 0x95, Template: "resnr", Bit: 2, Register: "L"},
		{Extended: true, Opcode: 0x96, Template: "resnaddr", Bit: 2, Register: "HL"},
		{Extended: true, Opcode: 0x97, Template: "resnr", Bit: 2, Register: "A"},
		{Extended: true, Opcode: 0x98, Template: "resnr", Bit: 3, Register: "B"},
		{Extended: true, Opcode: 0x99, Template: "resnr", Bit: 3, Register: "C"},
		{Extended: true, Opcode: 0x9a, Template: "resnr", Bit: 3, Register: "D"},
		{Extended: true, Opcode: 0x9b, Template: "resnr", Bit: 3, Register: "E"},
		{Extended: true, Opcode: 0x9c, Template: "resnr", Bit: 3, Register: "H"},
		{Extended: true, Opcode: 0x9d, Template: "resnr", Bit: 3, Register: "L"},
		{Extended: true, Opcode: 0x9e, Template: "resnaddr", Bit: 3, Register: "HL"},
		{Extended: true, Opcode: 0x9f, Template: "resnr", Bit: 3, Register: "A"},
		{Extended: true, Opcode: 0xa0, Template: "resnr", Bit: 4, Register: "B"},
		{Extended: true, Opcode: 0xa1, Template: "resnr", Bit: 4, Register: "C"},
		{Extended: true, Opcode: 0xa2, Template: "resnr", Bit: 4, Register: "D"},
		{Extended: true, Opcode: 0xa3, Template: "resnr", Bit: 4, Register: "E"},
		{Extended: true, Opcode: 0xa4, Template: "resnr", Bit: 4, Register: "H"},
		{Extended: true, Opcode: 0xa5, Template: "resnr", Bit: 4, Register: "L"},
		{Extended: true, Opcode: 0xa6, Template: "resnaddr", Bit: 4, Register: "HL"},
		{Extended: true, Opcode: 0xa7, Template: "resnr", Bit: 4, Register: "A"},
		{Extended: true, Opcode: 0xa8, Template: "resnr", Bit: 5, Register: "B"},
		{Extended: true, Opcode: 0xa9, Template: "resnr", Bit: 5, Register: "C"},
		{Extended: true, Opcode: 0xaa, Template: "resnr", Bit: 5, Register: "D"},
		{Extended: true, Opcode: 0xab, Template: "resnr", Bit: 5, Register: "E"},
		{Extended: true, Opcode: 0xac, Template: "resnr", Bit: 5, Register: "H"},
		{Extended: true, Opcode: 0xad, Template: "resnr", Bit: 5, Register: "L"},
		{Extended: true, Opcode: 0xae, Template: "resnaddr", Bit: 5, Register: "HL"},
		{Extended: true, Opcode: 0xaf, Template: "resnr", Bit: 5, Register: "A"},
		{Extended: true, Opcode: 0xb0, Template: "resnr", Bit: 6, Register: "B"},
		{Extended: true, Opcode: 0xb1, Template: "resnr", Bit: 6, Register: "C"},
		{Extended: true, Opcode: 0xb2, Template: "resnr", Bit: 6, Register: "D"},
		{Extended: true, Opcode: 0xb3, Template: "resnr", Bit: 6, Register: "E"},
		{Extended: true, Opcode: 0xb4, Template: "resnr", Bit: 6, Register: "H"},
		{Extended: true, Opcode: 0xb5, Template: "resnr", Bit: 6, Register: "L"},
		{Extended: true, Opcode: 0xb6, Template: "resnaddr", Bit: 6, Register: "HL"},
		{Extended: true, Opcode: 0xb7, Template: "resnr", Bit: 6, Register: "A"},
		{Extended: true, Opcode: 0xb8, Template: "resnr", Bit: 7, Register: "B"},
		{Extended: true, Opcode: 0xb9, Template: "resnr", Bit: 7, Register: "C"},
		{Extended: true, Opcode: 0xba, Template: "resnr", Bit: 7, Register: "D"},
		{Extended: true, Opcode: 0xbb, Template: "resnr", Bit: 7, Register: "E"},
		{Extended: true, Opcode: 0xbc, Template: "resnr", Bit: 7, Register: "H"},
		{Extended: true, Opcode: 0xbd, Template: "resnr", Bit: 7, Register: "L"},
		{Extended: true, Opcode: 0xbe, Template: "resnaddr", Bit: 7, Register: "HL"},
		{Extended: true, Opcode: 0xbf, Template: "resnr", Bit: 7, Register: "A"},
	}

	funcs := template.FuncMap{"date": time.Now, "lower": strings.ToLower, "name": typeName}
	t := template.New("header.gotmpl").Funcs(funcs)
	t = template.Must(t.ParseGlob("instructions/*.gotmpl"))

	// Render all instructions in opcode order.
	if f, err := os.Create("instructionset.go"); err == nil {
		defer func() {
			f.Close()
		}()

		// Header and instruction set arrays.
		err = t.Execute(f, instructionSet{instructions, extended})
		if err != nil {
			panic(err)
		}

		for _, i := range instructions {
			err = t.ExecuteTemplate(f, fmt.Sprintf("%s.gotmpl", i.Template), &i)
			if err != nil {
				panic(err)
			}
		}

		for _, i := range extended {
			err = t.ExecuteTemplate(f, fmt.Sprintf("%s.gotmpl", i.Template), &i)
			if err != nil {
				panic(err)
			}
		}
	}
}
