// Package logger allows configurable, per-package logging.
package logger

import (
	"fmt"
)

// Copy of the last string displayed to gracefully handle repeated text.
// TODO: making this into a FIFO with a configurable size would be awesome.
var lastMessage string

// Count of how many times the last message was repeated.
var lastMessageCount uint

// LogLevel representing the priority of a log message. An attempt to log a
// message whose LogLevel is above logger.Level will be silently ignored.
type LogLevel uint8

// Supported log levels.
const (
	Fatal LogLevel = iota
	Warning
	Info
	Debug
	Desperate // The kind of level used in a Pixel FIFO, for instance...
)

// Level is the global log level above which nothing will be displayed.
var Level = Info // Sensible default

// Enabled setting controls whether logging will occur for a given module name
// when the usual methods are called. Default is no logging. Enabling 'all'
// will turn on debug output for every module.
var Enabled = make(map[string]bool)

// Context is a function returning a string that will be prepended to every
// log message, if defined. This can be used to insert timestamps, CPU info...
// The default does nothing and debug messages will be displayed as given.
var Context = func() string { return "" }

// Logger represents a name-based logger for "simple" debugging. Each package
// should register itself with logger, which will allow for a complete listing
// of all possible logger package/modules that can be used with the -debug
// flag.
type Logger struct {
	pkg     string          // Package to which this logger applies.
	modules map[string]bool // Set of sub-categories this object supports.
	// For quick lookup later, they're stored as
	// <pkg name>/<module name> directly.

	pkgFull string // Cached string '<pkg name>/*' for quick lookup.
}

// Output log message if the given package/subpackage is enabled and if the
// global log level permits it.
func (l *Logger) log(level LogLevel, module, format string, a ...interface{}) {
	// Do we need to log this?
	if level > Level {
		return
	}

	enabled := false
	switch {
	case Enabled["all"]:
		enabled = true
	case module == "" && Enabled[l.pkg]:
		enabled = true
	case Enabled[module] || Enabled[l.pkgFull]:
		enabled = true
	}

	if !enabled {
		return
	}

	fmt.Print(Context())
	msg := fmt.Sprintf(format, a...)
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

// Fatal prints a message (then panics regardless of debug level).
func (l *Logger) Fatal(msg string) {
	l.log(Fatal, "", "%s", msg)
	panic(msg)
}

// Warning prints a message if the global log level is Warning or more.
func (l *Logger) Warning(msg string) {
	l.log(Warning, "", "%s", msg)
}

// Info prints a message if the global log level is Info (the default) or more.
func (l *Logger) Info(msg string) {
	l.log(Info, "", "%s", msg)
}

// Log is an alias for Info.
func (l *Logger) Log(msg string) {
	l.log(Info, "", "%s", msg)
}

// Debug prints a message if the global log level is Debug or more.
func (l *Logger) Debug(msg string) {
	l.log(Debug, "", "%s", msg)
}

// Desperate prints a message if the global log level is the maximum.
func (l *Logger) Desperate(msg string) {
	l.log(Desperate, "", "%s", msg)
}
