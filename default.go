package gocolorful

var defaultLevels = map[string]Level{
	"INFO": Level{
		"INFO",
		Blue,
	},
	"WARNING": Level{
		"WARNING",
		Yellow,
	},
	"DEBUG": Level{
		"DEBUG",
		White,
	},
	"ERR": Level{
		"ERR",
		Red,
	},
}

type output interface {
	Print(v ...interface{})
}

// Info print information to terminal
func Info(message ...interface{}) {
	defaultLevels["INFO"].Print(message...)
}

// Warning print warning to terminal
func Warning(message ...interface{}) {
	defaultLevels["WARNING"].Print(message...)
}

// Err print error to terminal
func Err(message ...interface{}) {
	defaultLevels["ERR"].Print(message...)
}

// Debug print debug to terminal
func Debug(message ...interface{}) {
	defaultLevels["DEBUG"].Print(message...)
}
