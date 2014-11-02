package golem

import (
	"github.com/mgutz/ansi"
	"log"
	"os"
)

var (
	i = ansi.ColorCode("white+b:blue")
	w = ansi.ColorCode("white+b:red")
	v = ansi.ColorCode("white+b:black")
	r = ansi.ColorCode("reset")
)

var (
	INFO = log.New(os.Stdout, i+"[INFO]"+r+" ", 0)
	WARN = log.New(os.Stdout, w+"[WARN]"+r+" ", 0)
	VERB = log.New(os.Stdout, v+"[VERB]"+r+" ", 0)
)

// Print info statements
func Info(i interface{}) {
	INFO.Print(i)
}

// Printf info statements
func Infof(s string, i interface{}) {
	INFO.Printf(s, i)
}

// Print warning statements
func Warn(i interface{}) {
	WARN.Print(i)
}

// Printf warning statements
func Warnf(s string, i interface{}) {
	WARN.Printf(s, i)
}

// Print verbose statements if debug is true
func Verb(i interface{}) {
	VERB.Print(i)
}

// Printf verbose statements if debug is true
func Verbf(s string, i interface{}) {
	VERB.Printf(s, i)
}
