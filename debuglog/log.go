package debuglog

import "log"

var Enabled = false

func Debugf(format string, v ...interface{}) {
	if !Enabled {
		return
	}

	log.Printf(format, v...)
}
