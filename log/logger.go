package log

import "fmt"

// Simplistic on/off logger for simple debugging purposes.

// Enabled setting controls whether logging will occur when the usual methods are called. Default is no logging.
var Enabled = false

// Print displays a message just like fmt.Print but only if Enabled is true.
func Print(v ...interface{}) {
	if Enabled {
		fmt.Print(v...)
	}
}

// Printf displays a message just like fmt.Printf but only if Enabled is true.
func Printf(format string, v ...interface{}) {
	if Enabled {
		fmt.Printf(format, v...)
	}
}

// Println displays a message just like fmt.Println but only if Enabled is true.
func Println(v ...interface{}) {
	if Enabled {
		fmt.Println(v...)
	}
}
