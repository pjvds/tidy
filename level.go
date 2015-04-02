package tidy

const (
	FATAL Level = iota
	ERROR
	WARN
	NOTICE
	INFO
	DEBUG
)

var (
	chars = [...]string{
		string(names[FATAL][0]),
		string(names[ERROR][0]),
		string(names[WARN][0]),
		string(names[NOTICE][0]),
		string(names[INFO][0]),
		string(names[DEBUG][0]),
	}
	names = [...]string{
		"FATAL",
		"ERROR",
		"WARN",
		"NOTICE",
		"INFO",
		"DEBUG",
	}
)

// Represents a serverity level. The levels are:
// FATAL, ERROR, WARN, NOTICE, INFO and DEBUG.
// Use the Allows method to see if the current level
// allows a specified level, for example:
//
//     DEBUG.Allows(ERROR) => true
type Level byte

// Allows returns true if the level allows the specified level;
// otherwise false. An level always allows itself and every other
// level that is above it. INFO would allow INFO and WARN, but
// it would not allow DEBUG.
func (this Level) Allows(other Level) bool {
	return this >= other
}

// String returns the level fullname, like: FATAL, ERROR or DEBUG.
func (this Level) String() string {
	return names[this]
}
