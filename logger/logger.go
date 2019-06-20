// Package logger allows configurable, per-package logging.
package logger

import (
	"bytes"
	"fmt"
	"sort"
)

// Copy of the last string displayed to gracefully handle repeated text.
// TODO: making this into a FIFO with a configurable size would be awesome.
//       (I now regret moving fifo.go to the ppu package...)
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

// Loggers is a registry of currently defined package-specific loggers.
var Loggers = make(map[string]*Logger)

// Context is a function returning a string that will be prepended to every
// log message, if defined. This can be used to insert timestamps, CPU info...
// The default does nothing and debug messages will be displayed as given.
var Context = func() string { return "" }

// Logger represents a name-based logger for "simple" debugging. Each package
// should register itself with logger, which will allow for a complete listing
// of all possible logger package/modules that can be used with the -debug
// flag.
type Logger struct {
	Name string // Package or module name for this logger
	Help string // Help text to be displayed when calling with -debug help

	wildcard string             // Cached '<pkg name>/*' for quick lookup.
	modules  map[string]*Logger // Submodules (nil if Logger is already one)
}

// New returns a Logger instance specific to the given package, after
// registering it with our base logger package so this logger and its modules
// can be listed from the command-line.
// Will panic if a logger with the same name is already defined.
func New(name, help string) *Logger {
	if Loggers[name] != nil {
		panic(fmt.Sprintf("logger '%s' already exists", name))
	}
	l := &Logger{name, help, fmt.Sprintf("%s/*", name), make(map[string]*Logger)}
	// Record this so we have exactly one logger per package.
	Loggers[name] = l
	return l
}

// Add a sub-module to a package logger, which can then be enabled by using
// -debug <pkg>/<module> or -debug <pkg>/* on the command line.
func (l *Logger) Add(name, help string) {
	modName := fmt.Sprintf("%s/%s", l.Name, name)
	if l.modules[name] != nil {
		panic(fmt.Sprintf("logger '%s' already exists", modName))
	}
	l.modules[name] = &Logger{modName, help, l.wildcard, nil}
}

// Sub returns a sub-module with the given name. If called with a module name
// that hasn't been registered with Add() first, returns the main logger.
// Otherwise, returns a Logger instance to allow chain call to print methods.
func (l *Logger) Sub(name string) *Logger {
	if sub := l.modules[name]; sub != nil {
		return sub
	}
	fmt.Printf(" !!! sub-logger %s/%s not found\n", l.Name, name)
	return l
}

// Output log message if the given package/subpackage is enabled and if the
// global log level permits it.
func (l *Logger) log(level LogLevel, format string, a ...interface{}) {
	// "Do we need to log this?"
	if level > Level {
		return
	}

	if !(Enabled["all"] || Enabled[l.Name] || Enabled[l.wildcard]) {
		return
	}

	fmt.Print(Context())
	msg := fmt.Sprintf("%s: %s", l.Name, fmt.Sprintf(format, a...))
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
	l.log(Fatal, "%s", msg)
	panic(msg)
}

// Warning prints a message if the global log level is Warning or more.
func (l *Logger) Warning(msg string) {
	l.log(Warning, "%s", msg)
}

// Info prints a message if the global log level is Info (the default) or more.
func (l *Logger) Info(msg string) {
	l.log(Info, "%s", msg)
}

// Log is an alias for Info.
func (l *Logger) Log(msg string) {
	l.log(Info, "%s", msg)
}

// Debug prints a message if the global log level is Debug or more.
func (l *Logger) Debug(msg string) {
	l.log(Debug, "%s", msg)
}

// Desperate prints a message if the global log level is the maximum.
func (l *Logger) Desperate(msg string) {
	l.log(Desperate, "%s", msg)
}

// Fatalf format-prints a message (then panics regardless of debug level).
func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.log(Fatal, format, a...)
	panic(fmt.Sprintf(format, a...))
}

// Warningf format-prints a message if the global log level is Warning or more.
func (l *Logger) Warningf(format string, a ...interface{}) {
	l.log(Warning, format, a...)
}

// Infof format-prints a message if the global log level is Info (the default) or more.
func (l *Logger) Infof(format string, a ...interface{}) {
	l.log(Info, format, a...)
}

// Logf is an alias for Infof.
func (l *Logger) Logf(format string, a ...interface{}) {
	l.log(Info, format, a...)
}

// Debugf format-prints a message if the global log level is Debug or more.
func (l *Logger) Debugf(format string, a ...interface{}) {
	l.log(Debug, format, a...)
}

// Desperatef format-prints a message if the global log level is the maximum.
func (l *Logger) Desperatef(format string, a ...interface{}) {
	l.log(Desperate, format, a...)
}

// String returns the description of a logger and all its sub-loggers if any.
func (l *Logger) String() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "%s: %s\n", l.Name, l.Help)
	if l.modules != nil {
		names := make([]string, 0, len(l.modules))
		for n := range l.modules {
			names = append(names, n)
		}
		sort.Strings(names)
		for _, n := range names {
			fmt.Fprint(&b, l.modules[n])
		}
	}
	return b.String()
}

// Help prints all registered loggers and sub-loggers so the user knows what
// can be enabled.
func Help() {
	// Sort loggers by package name.
	names := make([]string, 0, len(Loggers))
	for n := range Loggers {
		names = append(names, n)
	}
	sort.Strings(names)

	// Display all loggers and their submodules (if any, also sorted).
	for _, n := range names {
		fmt.Println(Loggers[n])
	}
}
