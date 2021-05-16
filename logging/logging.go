package logging

import (
	"fmt"
	"log"
)

type Logger struct {
	Verbose bool
}

func (l Logger) Debug(format string, v ... interface{}) {
	if l.Verbose {
		log.Printf(format, v ...)
	}
}

func (l Logger) Info(format string, v ... interface{}) {
	fmt.Printf(format, v ...)
}
