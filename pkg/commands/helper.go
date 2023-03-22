package commands

import "github.com/fatih/color"

// bad api and abstraction

func green(format string, args ...interface{}) {
	_, _ = color.New(color.FgGreen).Fprintf(color.Error, format, args...)
}

func red(format string, args ...interface{}) {
	_, _ = color.New(color.FgRed).Fprintf(color.Error, format, args...)
}
