package logging

import (
	"path/filepath"
	"runtime"
	"strconv"
)

type Location string

func GetLocation(depth int) Location {
	_, file, line, ok := runtime.Caller(1 + depth)

	if !ok {
		return Location("")
	}

	return Location(filepath.Base(file) + ":" + strconv.Itoa(line))
}

func (this Location) IsEmpty() bool {
	return len(this) == 0
}

func (this Location) String() string {
	return string(this)
}
