package config

import (
	"fmt"
)

const (
	//Black is a constant of type int
	Black = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

var (
	//InfoByte is a variable of type []byte
	DebugByteColor = []byte(fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", Blue, "DEBUG"))
	WarnByteColor  = []byte(fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", Yellow, "WARN"))
	ErrorByteColor = []byte(fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", Red, "ERROR"))
	FatalByteColor = []byte(fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", Magenta, "FATAL"))
)
