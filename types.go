package gocolorful

// Color defined ANSI-ESCAPE sequence char for color
var Color = map[string]string{
	"Purple": "\u001b[95m",
	"Blue":   "\u001b[94m",
	"Yellow": "\u001b[93m",
	"Green":  "\u001b[92m",
	"Red":    "\u001b[91m",
}

// BrightColor defined ANSI-ESCAPE sequence char for bright color
var BrightColor = map[string]string{
	"Black":   "\u001b[40;1m",
	"Red":     "\u001b[41;1m",
	"Green":   "\u001b[42;1m",
	"Yellow":  "\u001b[43;1m",
	"Blue":    "\u001b[44;1m",
	"Magenta": "\u001b[45;1m",
	"Cyan":    "\u001b[46;1m",
	"White":   "\u001b[47;1m",
}

// End means ANSI-ESCAPE sequence end.
// Should be Placed in each line end or new line start.
var End = "\u001b[0m"

// Underline defined a special type of underline represent
var Underline = "\u001b[4m"

// Bold defined bold text format
var Bold = "\033[1m"
