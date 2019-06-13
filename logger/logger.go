// Package logger allows configurable, per-package logging.
package logger

import (
	"fmt"
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
	pkg string // Package to which this logger applies.

	pkgFull string            // Cached '<pkg name>/*' for quick lookup.
	modules map[string]string // Cached '<pkg name>/<module name>'.
}

// New returns a Logger instance specific to the given package, after
// registering it with our base logger package so this logger and its modules
// can be listed from the command-line.
// TODO: add a mandatory description parameter. And for modules too.
func New(pkg string, modules []string) (*Logger, error) {
	if Loggers[pkg] != nil {
		return nil, fmt.Errorf("package %s already has a logger", pkg)
	}

	l := Logger{pkg, fmt.Sprintf("%s/*", pkg), make(map[string]string)}
	for _, m := range modules {
		l.modules[m] = fmt.Sprintf("%s/%s", pkg, m)
	}

	// Record this so we have exactly one logger per package.
	Loggers[pkg] = &l

	return &l, nil
}

// Output log message if the given package/subpackage is enabled and if the
// global log level permits it.
func (l *Logger) log(level LogLevel, module, format string, a ...interface{}) {
	// "Do we need to log this?"
	if level > Level {
		return
	}
	moduleFull := fmt.Sprintf("%s/%s", l.pkg, module)
	enabled := false
	switch {
	case Enabled["all"]:
		enabled = true
	case module == "" && Enabled[l.pkg]:
		enabled = true
	case Enabled[moduleFull] || Enabled[l.pkgFull]:
		enabled = true
	}

	if !enabled {
		return
	}

	fmt.Printf("%s%s: ", Context(), moduleFull)
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

//
// Package-only logging functions.
//

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

// Fatalf format-prints a message (then panics regardless of debug level).
func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.log(Fatal, "", format, a...)
	panic(fmt.Sprintf(format, a...))
}

// Warningf format-prints a message if the global log level is Warning or more.
func (l *Logger) Warningf(format string, a ...interface{}) {
	l.log(Warning, "", format, a...)
}

// Infof format-prints a message if the global log level is Info (the default) or more.
func (l *Logger) Infof(format string, a ...interface{}) {
	l.log(Info, "", format, a...)
}

// Logf is an alias for Infof.
func (l *Logger) Logf(format string, a ...interface{}) {
	l.log(Info, "", format, a...)
}

// Debugf format-prints a message if the global log level is Debug or more.
func (l *Logger) Debugf(format string, a ...interface{}) {
	l.log(Debug, "", format, a...)
}

// Desperatef format-prints a message if the global log level is the maximum.
func (l *Logger) Desperatef(format string, a ...interface{}) {
	l.log(Desperate, "", format, a...)
}

//
// Sub-module logging functions.
//

// FatalMod prints a message for the given module (then panics regardless of
// debug level).
func (l *Logger) FatalMod(module, msg string) {
	l.log(Fatal, module, "%s", msg)
	panic(msg)
}

// WarningMod prints a message for the given module if the global log level is
// Warning or more.
func (l *Logger) WarningMod(module, msg string) {
	l.log(Warning, module, "%s", msg)
}

// InfoMod prints a message for the given module if the global log level is Info
// (the default) or more.
func (l *Logger) InfoMod(module, msg string) {
	l.log(Info, module, "%s", msg)
}

// LogMod is an alias for InfoMod.
func (l *Logger) LogMod(module, msg string) {
	l.log(Info, module, "%s", msg)
}

// DebugMod prints a message for the given module if the global log level is
// Debug or more.
func (l *Logger) DebugMod(module, msg string) {
	l.log(Debug, module, "%s", msg)
}

// DesperateMod prints a message for the given module if the global log level is
// the maximum.
func (l *Logger) DesperateMod(module, msg string) {
	l.log(Desperate, module, "%s", msg)
}

// FatalfMod format-prints a message for the given module (then panics
// regardless of debug level).
func (l *Logger) FatalfMod(module, format string, a ...interface{}) {
	l.log(Fatal, module, format, a...)
	panic(fmt.Sprintf(format, a...))
}

// WarningfMod format-prints a message if the global log level is Warning or more.
func (l *Logger) WarningfMod(module, format string, a ...interface{}) {
	l.log(Warning, module, format, a...)
}

// InfofMod format-prints a message if the global log level is Info (the default) or more.
func (l *Logger) InfofMod(module, format string, a ...interface{}) {
	l.log(Info, module, format, a...)
}

// LogfMod is an alias for InfofMod.
func (l *Logger) LogfMod(module, format string, a ...interface{}) {
	l.log(Info, module, format, a...)
}

// DebugfMod format-prints a message if the global log level is Debug or more.
func (l *Logger) DebugfMod(module, format string, a ...interface{}) {
	l.log(Debug, module, format, a...)
}

// DesperatefMod format-prints a message if the global log level is the maximum.
func (l *Logger) DesperatefMod(module, format string, a ...interface{}) {
	l.log(Desperate, module, format, a...)
}
