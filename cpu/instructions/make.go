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

func main() {

	instructions := []data{
		{Opcode: 0x00, Template: "nop"},
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

	funcs := template.FuncMap{"lower": strings.ToLower}
	t := template.New("header.tmpl").Funcs(funcs)
	t = template.Must(t.ParseGlob("instructions/*.tmpl"))

	// Render all instructions in opcode order.
	if f, err := os.Create("instructionset.go"); err == nil {
		defer func() {
			f.Close()
		}()

		t.Execute(f, nil)

		for _, i := range instructions {
			t.ExecuteTemplate(f, fmt.Sprintf("%s.tmpl", i.Template), i)
		}
	}
}
