package log

import (
	"log"
)

// Name-based on/off logger for simple debugging purposes.

// Enabled setting controls whether logging will occur for a given module name when the usual methods are called. Default is no logging.
var Enabled = make(map[string]bool)

// Print displays a message just like fmt.Print but only if Enabled is true.
func Print(name string, v ...interface{}) {
	if Enabled[name] {
		log.Print(v...)
	}
}

// Printf displays a message just like fmt.Printf but only if Enabled is true.
func Printf(name string, format string, v ...interface{}) {
	if Enabled[name] {
		log.Printf(format, v...)
	}
}

// Println displays a message just like fmt.Println but only if Enabled is true.
func Println(name string, v ...interface{}) {
	if Enabled[name] {
		log.Println(v...)
	}
}
