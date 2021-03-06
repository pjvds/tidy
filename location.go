package tidy

import (
	"path/filepath"
	"runtime"
	"strconv"
)

// Represents the location that wrote the log entry.
// A location is holds the base filename and line number
// in the following format:
//
//     file.go:42
type Location string

// GetLocation returns the location of the caller.
func GetLocation(depth int) Location {
	var location Location

	if _, file, line, ok := runtime.Caller(1 + depth); ok {
		location = Location(filepath.Base(file) + ":" + strconv.Itoa(line))
	}

	return location
}

func (this Location) IsEmpty() bool {
	return len(this) == 0
}

func (this Location) String() string {
	return string(this)
}
