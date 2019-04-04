package debug

import (
	"fmt"
)

// Name-based on/off logger for simple debugging purposes.

// Enabled setting controls whether logging will occur for a given module name
// when the usual methods are called. Default is no logging. Enabling 'all'
// will turn on debug output for every module.
var Enabled = make(map[string]bool)

// Context is a function returning a string that will be prepended to every
// log message, if defined. This can be used to insert timestamps, CPU info...
// The default does nothing and debug messages will be displayed as given.
var Context = func() string { return "" }

// Print displays a message just like fmt.Print but only if Enabled is true
// for the given module name (or if 'all' was enabled). Also it automatically
// tacks a linefeed at the end.
func Print(name string, v ...interface{}) {
	if Enabled[name] || Enabled["all"] {
		fmt.Print(Context())
		fmt.Println(fmt.Sprint(v...))
	}
}

// Printf displays a message just like fmt.Printf but only if Enabled is true
// for the given module name (or if 'all' was enabled). Also it automatically
// tacks a linefeed at the end.
func Printf(name string, format string, v ...interface{}) {
	if Enabled[name] || Enabled["all"] {
		fmt.Print(Context())
		fmt.Println(fmt.Sprintf(format, v...))
	}
}

// Println displays a message just like fmt.Println but only if Enabled is true
// for the given module name (or if 'all' was enabled).
func Println(name string, v ...interface{}) {
	if Enabled[name] || Enabled["all"] {
		fmt.Print(Context())
		fmt.Println(v...)
	}
}
