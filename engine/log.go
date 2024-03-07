package engine

import (
	"fmt"
)

func (e *Engine) logDebug(format string, a ...any) (n int, err error) {
	if !e.debug {
		return
	}

	return fmt.Fprintf(e.logWriter, format+"\n", a...)
}

func (e *Engine) log(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(e.logWriter, format+"\n", a...)
}
