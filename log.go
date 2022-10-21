package loglevel

import (
	"fmt"
	"strings"

	"github.com/Masterminds/log-go"
)

// ParseLevel returns the loglevel integer associated with a common loglevel
// string representation.
func ParseLevel(loglevelStr string) (level int, err error) {
	switch strings.ToLower(loglevelStr) {
	case "trace":
		level = log.TraceLevel
	case "debug":
		level = log.DebugLevel
	case "info":
		level = log.InfoLevel
	case "warn":
		level = log.WarnLevel
	case "warning":
		level = log.WarnLevel
	case "error":
		level = log.ErrorLevel
	case "panic":
		level = log.PanicLevel
	case "fatal":
		level = log.FatalLevel
	default:
		err = fmt.Errorf("unknown loglevel: %s", loglevelStr)
	}
	return
}
