package log

import (
	"net/http"
	"time"
)

type MsgType uint8
type OutType uint8

const (
	MessageLog         MsgType = 0
	Message2Log        MsgType = 1
	WarningLog         MsgType = 2
	DebugLog           MsgType = 3
	ErrorLog           MsgType = 4
	FormattedOut       OutType = 0
	LineOut            OutType = 1
	DefaultMaxLineSize int     = 2000
)

var (
	// DebugMode Enable debug mode
	DebugMode bool

	// MaxLineSize limits the size of the line, if the size
	// exceeds that indicated by MaxLineSize the system cuts
	// the string and adds "..." at the end.
	MaxLineSize = DefaultMaxLineSize

	// Colors contain color array
	Colors = []string{
		MessageLog:  "\x1b[37m", // White
		Message2Log: "\x1b[92m", // Light green
		WarningLog:  "\x1b[93m", // Light Yellow
		DebugLog:    "\x1b[96m", // Light Cyan
		ErrorLog:    "\x1b[91m", // Light Red
	}

	// Prefixes of messages
	Prefixes = []string{
		MessageLog:  "msg",
		Message2Log: "msg",
		WarningLog:  "warning",
		DebugLog:    "debug",
		ErrorLog:    "error",
	}

	Now = time.Now

	Loggers []Logger
)

type Logger struct {
	HTTPError func(w http.ResponseWriter, code int)
	Fatal     func(msg ...interface{})
	Errorln   func(msg ...interface{})
	Errorf    func(msg ...interface{})
	Warningln func(msg ...interface{})
	Warningf  func(msg ...interface{})
	Println   func(msg ...interface{})
	Printf    func(msg ...interface{})
	Debugln   func(msg ...interface{})
	Debugf    func(msg ...interface{})
}

// HTTPError write lot to stdout and return json error on http.ResponseWriter with http error code.
func HTTPError(w http.ResponseWriter, code int) {
	for _, l := range Loggers {
		l.HTTPError(w, code)
	}
}

// Fatal show message with line break at the end and exit to OS.
func Fatal(msg ...interface{}) {
	for _, l := range Loggers {
		l.Fatal(msg)
	}
}

// Errorln message with line break at the end.
func Errorln(msg ...interface{}) {
	for _, l := range Loggers {
		l.Errorln(msg)
	}
}

// Errorf shows formatted error message on stdout without line break at the end.
func Errorf(msg ...interface{}) {
	for _, l := range Loggers {
		l.Errorf(msg)
	}
}

// Warningln shows warning message on stdout with line break at the end.
func Warningln(msg ...interface{}) {
	for _, l := range Loggers {
		l.Warningln(msg)
	}
}

// Warningf shows formatted warning message on stdout without line break at the end.
func Warningf(msg ...interface{}) {
	for _, l := range Loggers {
		l.Warningf(msg)
	}
}

// Println shows message on stdout with line break at the end.
func Println(msg ...interface{}) {
	for _, l := range Loggers {
		l.Println(msg)
	}
}

// Printf shows formatted message on stdout without line break at the end.
func Printf(msg ...interface{}) {
	for _, l := range Loggers {
		l.Printf(msg)
	}
}

// Debugln shows debug message on stdout with line break at the end.
// If debug mode is not active no message is displayed
func Debugln(msg ...interface{}) {
	for _, l := range Loggers {
		l.Debugln(msg)
	}
}

// Debugf shows debug message on stdout without line break at the end.
// If debug mode is not active no message is displayed
func Debugf(msg ...interface{}) {
	for _, l := range Loggers {
		l.Debugf(msg)
	}
}
