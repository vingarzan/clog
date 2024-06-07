// Go SIP stack by Neat Path Networks
// Copyright ©️ 2024 Neat Path Networks GmbH
// Authors(s):
//   - Dragos Vingarzan - dragos@neatpath.net

package clog

import "io"

// Logging - copy this to your package and adjust the name

var (
	lDebug   *CustomLogger
	lInfo    *CustomLogger
	lWarning *CustomLogger
	lError   *CustomLogger
)

const packageName = "custom-log"

func init() {
	lDebug, lInfo, lWarning, lError = MakeLoggers(packageName)
}

func LogInit(out io.Writer) {
	lDebug, lInfo, lWarning, lError = MakeLoggersWithOutput(packageName, out)
}

// package yourpackage
//
// import clog "github.com/vingarzan/clog"

// var (
// 	lDebug   *clog.CustomLogger
// 	lInfo    *clog.CustomLogger
// 	lWarning *clog.CustomLogger
// 	lError   *clog.CustomLogger
// )

// const packageName = "your-package-name" // keep it short, at most 12 characters for best results

// func init() {
// 	lDebug, lInfo, lWarning, lError = clog.MakeLoggers(packageName)
// }

// func LogInit(out io.Writer) {
// 	lDebug, lInfo, lWarning, lError = clog.MakeLoggersWithOutput(packageName, out)
// }
