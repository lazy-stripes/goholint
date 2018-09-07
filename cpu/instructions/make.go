package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// Quick and dirty aggregate of keys, not necessarily all used by all instructions.
type data struct {
	Opcode      uint
	Template    string
	Instruction string
	High, Low   string
	Register    string
	Address     string
	Operator    string
}

// typeName generates a deterministic name for derived instruction types.
// TODO: I'm going the easy way for now, but I'd love being able to generate
//		 human-readable name like nop, ldDeD16, etc.
func typeName(d *data) (name string) {
	return fmt.Sprintf("op%s", strings.Title(fmt.Sprintf("%02x", d.Opcode)))
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

		{Opcode: 0x0c, Template: "incr", Register: "C"},
		{Opcode: 0x0d, Template: "decr", Register: "C"},
		{Opcode: 0x0e, Template: "ldrd8", Register: "C"},
		{Opcode: 0x11, Template: "ldrrd16", High: "D", Low: "E"},
		{Opcode: 0x13, Template: "incrr", High: "D", Low: "E"},
		{Opcode: 0x14, Template: "incr", Register: "D"},
		{Opcode: 0x15, Template: "decr", Register: "D"},
		{Opcode: 0x16, Template: "ldrd8", Register: "D"},
		{Opcode: 0x1c, Template: "incr", Register: "E"},
		{Opcode: 0x1d, Template: "decr", Register: "E"},
		{Opcode: 0x1e, Template: "ldrd8", Register: "E"},
		{Opcode: 0x21, Template: "ldrrd16", High: "H", Low: "L"},
		{Opcode: 0x23, Template: "incrr", High: "H", Low: "L"},
		{Opcode: 0x24, Template: "incr", Register: "H"},
		{Opcode: 0x25, Template: "decr", Register: "H"},
		{Opcode: 0x26, Template: "ldrd8", Register: "H"},
		{Opcode: 0x2c, Template: "incr", Register: "L"},
		{Opcode: 0x2d, Template: "decr", Register: "L"},
		{Opcode: 0x2e, Template: "ldrd8", Register: "L"},
		{Opcode: 0x3c, Template: "incr", Register: "A"},
		{Opcode: 0x3d, Template: "decr", Register: "A"},
		{Opcode: 0x3e, Template: "ldrd8", Register: "A"},
		{Opcode: 0xa0, Template: "boolreg", Instruction: "AND", Operator: "&=", Register: "B"},
		{Opcode: 0xa1, Template: "boolreg", Instruction: "AND", Operator: "&=", Register: "C"},
		{Opcode: 0xa2, Template: "boolreg", Instruction: "AND", Operator: "&=", Register: "D"},
		{Opcode: 0xa3, Template: "boolreg", Instruction: "AND", Operator: "&=", Register: "E"},
		{Opcode: 0xa4, Template: "boolreg", Instruction: "AND", Operator: "&=", Register: "H"},
		{Opcode: 0xa5, Template: "boolreg", Instruction: "AND", Operator: "&=", Register: "L"},
		{Opcode: 0xa6, Template: "booladdr", Instruction: "AND", Operator: "&=", Address: "HL"},

		// 0xa6: andAddrHl,
		// 0xa7: andA,
		// 0xa8: xorB,
		// 0xa9: xorC,
		// 0xaa: xorD,
		// 0xab: xorE,
		// 0xac: xorH,
		// 0xad: xorL,
		// 0xae: xorAddrHl,
		// 0xaf: xorA,
		// 0xb0: orB,
		// 0xb1: orC,
		// 0xb2: orD,
		// 0xb3: orE,
		// 0xb4: orH,
		// 0xb5: orL,
		// 0xb6: orAddrHl,
		// 0xb7: orA,
	}

	funcs := template.FuncMap{"lower": strings.ToLower, "name": typeName}
	t := template.New("header.gotmpl").Funcs(funcs)
	t = template.Must(t.ParseGlob("instructions/*.gotmpl"))

	// Render all instructions in opcode order.
	if f, err := os.Create("instructionset.go"); err == nil {
		defer func() {
			f.Close()
		}()

		// Header and instruction set array.
		t.Execute(f, instructions)

		for _, i := range instructions {
			err = t.ExecuteTemplate(f, fmt.Sprintf("%s.gotmpl", i.Template), &i)
			if err != nil {
				panic(err)
			}
		}
	}
}
