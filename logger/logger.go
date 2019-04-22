// Package logger allows configurable, per-package logging.
package logger

import (
	"fmt"
)

// Name-based on/off logger for simple debugging purposes.

// Copy of the last string displayed to gracefully handle repeated text.
// TODO: making this into a FIFO with a configurable size would be awesome.
var lastMessage string

// Count of how many times the last message was repeated.
var lastMessageCount uint

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
		msg := fmt.Sprint(v...)
		if msg == lastMessage {
			lastMessageCount++
			fmt.Printf(" ... repeated %d times\r", lastMessageCount)
		} else {
			if lastMessageCount > 1 {
				fmt.Println()
			}
			lastMessage = msg
			lastMessageCount = 1
			fmt.Println(msg)
		}
	}
}

// Printf displays a message just like fmt.Printf but only if Enabled is true
// for the given module name (or if 'all' was enabled). Also it automatically
// tacks a linefeed at the end.
func Printf(name string, format string, v ...interface{}) {
	Print(name, fmt.Sprintf(format, v...))
}

// Println displays a message just like fmt.Println but only if Enabled is true
// for the given module name (or if 'all' was enabled).
func Println(name string, v ...interface{}) {
	Print(name, v...)
}
