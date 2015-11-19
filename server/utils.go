package server

import (
	"log"
)

func logf(format string, values ...interface{}) {
	if Verbose {
		log.Printf(format, values...)
	}
}
