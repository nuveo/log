package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

type msgType uint8

const (
	MessageLog  msgType = 0
	Message2Log msgType = 1
	WarningLog  msgType = 2
	DebugLog    msgType = 3
	ErrorLog    msgType = 4
)

var (

	// DebugMode Enable debug mode
	DebugMode bool

	Colors = []string{
		MessageLog:  "\x1b[37m", // White
		Message2Log: "\x1b[92m", // Light green
		WarningLog:  "\x1b[93m", // Light Yellow
		DebugLog:    "\x1b[96m", // Light Cyan
		ErrorLog:    "\x1b[91m", // Light Red
	}

	Prefixes = []string{
		MessageLog:  "msg",
		Message2Log: "msg",
		WarningLog:  "warning",
		DebugLog:    "debug",
		ErrorLog:    "error",
	}

	wg        sync.WaitGroup
	logChanel chan msgChanData
)

type msgChanData struct {
	mt   msgType
	msg  []interface{}
	line int
	file string
}

func init() {
	wg = sync.WaitGroup{}
	logChanel = make(chan msgChanData)

	go func() {
		for {
			m := <-logChanel
			pln(m)
			wg.Done()
		}
	}()
}

func Wait() {
	wg.Wait()
}

// Fatal show message with line break at the end and exit to OS.
func Fatal(msg ...interface{}) {
	send(ErrorLog, msg)
	os.Exit(-1)
}

// Errorln message with line break at the end.
func Errorln(msg ...interface{}) {
	send(ErrorLog, msg)
}

// Println shows message on screen with line break at the end.
func Println(msg ...interface{}) {
	send(MessageLog, msg)
}

// Debugln shows debug message on screen with line break at the end.
// If debug mode is not active no message is displayed
func Debugln(msg ...interface{}) {
	send(DebugLog, msg)
}

func send(mt msgType, msg ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	wg.Add(1)
	logChanel <- msgChanData{mt: mt, msg: msg, line: line, file: file}
}

func pln(m msgChanData) {
	if m.mt == DebugLog && !DebugMode {
		return
	}

	var debugInfo string

	if DebugMode {
		m.file = filepath.Base(m.file)
		debugInfo = fmt.Sprintf("%s:%d ", m.file, m.line)
	}

	fmt.Printf("%s%s [%s] %s%s\033[0;00m\n",
		Colors[m.mt],
		time.Now().UTC().Format("2006/01/02 15:04:05"),
		Prefixes[m.mt],
		debugInfo,
		fmt.Sprint(m.msg...))
}
