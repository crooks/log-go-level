package loglevel

import (
	"testing"

	"github.com/Masterminds/log-go"
)

func TestParseLevel(t *testing.T) {
	names := []string{"Trace", "Debug", "Info", "Warn", "Warning", "Error", "Panic", "Fatal"}
	levels := []int{
		log.TraceLevel,
		log.DebugLevel,
		log.InfoLevel,
		log.WarnLevel,
		log.WarnLevel,
		log.ErrorLevel,
		log.PanicLevel,
		log.FatalLevel,
	}
	if len(names) != len(levels) {
		t.Fatal("inconsistent number of elements in test")
	}
	for n := range names {
		name := names[n]
		level := levels[n]
		testLevel, err := ParseLevel(name)
		if err != nil {
			t.Errorf("name to level translation returned: %v", err)
		}
		if testLevel != level {
			t.Errorf("incorrect loglevel (%d) set by name \"%s\"", level, name)
		}
	}
	_, err := ParseLevel("fakeName")
	if err == nil {
		t.Error("fake loglevel name failed to return an error")
	}
}
