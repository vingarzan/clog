// Go SIP stack by Neat Path Networks
// Copyright ©️ 2024 Neat Path Networks GmbH
// Authors(s):
//   - Dragos Vingarzan - dragos@neatpath.net

package clog

// Logging

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

const (
	ANSI_RESET  = "\033[0m"
	ANSI_GRAY   = "\033[38;5;8m"
	ANSI_GREEN  = "\033[38;5;2m"
	ANSI_YELLOW = "\033[38;5;3m"
	ANSI_RED    = "\033[38;5;1m"
)

type CustomLogger struct {
	name   string
	logger *log.Logger
}

var spaces20 = []byte("                    ")

func (c *CustomLogger) getPrefix() string {
	_, file, line, _ := runtime.Caller(2)
	file = filepath.Base(file)

	b := bytes.NewBuffer(make([]byte, 0, 64))
	fmt.Fprintf(b, "| %s:%d", file, line)
	spaces := 24 - b.Len()
	if spaces > 0 {
		b.Write(spaces20[:spaces])
	}
	b.WriteString(c.name)
	return b.String()
}

func (c *CustomLogger) Print(v ...interface{}) {
	params := make([]interface{}, len(v)+2)
	params[0] = c.getPrefix()
	copy(params[1:], v)
	params[len(v)+1] = ANSI_RESET
	c.logger.Print(params...)
}

func (c *CustomLogger) Println(v ...interface{}) {
	params := make([]interface{}, len(v)+2)
	params[0] = c.getPrefix()
	copy(params[1:], v)
	params[len(v)+1] = ANSI_RESET
	c.logger.Println(params...)
}

func (c *CustomLogger) Printf(format string, v ...interface{}) {
	for i := len(format) - 1; i >= 0; i-- {
		if format[i] != '\n' && format[i] != '\r' {
			if i != len(format)-1 {
				format = format[:i+1]
			}
			break
		}
	}
	c.logger.Printf(c.getPrefix()+format+ANSI_RESET, v...)
}

func MakeLoggers(name string) (*CustomLogger, *CustomLogger, *CustomLogger, *CustomLogger) {

	return MakeLoggersWithOutput(name, nil)
}

func MakeLoggersWithOutput(name string, out io.Writer) (*CustomLogger, *CustomLogger, *CustomLogger, *CustomLogger) {

	// we could switch to log/slog for structured logging... that has built-in levels and attributes, context,
	// formatting, etc... but doesn't have the simpler Printf.

	// logFormat := log.Ldate | log.Lmicroseconds | log.Ltime | log.Lshortfile
	logFormat := log.Ldate | log.Lmicroseconds | log.Ltime
	// logFormat |= log.Lmsgprefix

	prefix := fmt.Sprintf("%-12s", name)
	prefix = prefix[:12]

	if out == nil {
		out = os.Stderr
	}

	debug := &CustomLogger{
		name:   "| " + prefix + " | DBG  | ",
		logger: log.New(out, ANSI_GRAY, logFormat),
	}
	info := &CustomLogger{
		name:   "| " + prefix + " | INFO | ",
		logger: log.New(out, "", logFormat),
	}
	warning := &CustomLogger{
		name:   "| " + prefix + " | WARN | ",
		logger: log.New(out, ANSI_YELLOW, logFormat),
	}
	err := &CustomLogger{
		name:   "| " + prefix + " | ERR  | ",
		logger: log.New(out, ANSI_RED, logFormat),
	}

	// setting also the default logger, just in case something uses just log.stuff
	log.SetFlags(logFormat)
	log.SetPrefix(prefix + "| ")

	return debug, info, warning, err
}
