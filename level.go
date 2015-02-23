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
	names = [...]string{
		"FATAL",
		"ERROR",
		"WARN",
		"NOTICE",
		"INFO",
		"DEBUG",
	}
	fixedNames = [...]string{
		"FATAL ",
		"ERROR ",
		"WARN  ",
		"NOTICE",
		"INFO  ",
		"DEBUG ",
	}
)

type Level byte

// Allows returns true if the level allows the specified level;
// otherwise false. An level always allows itself and every other
// level that is above it. INFO would allow INFO and WARN, but
// it would not allow DEBUG.
func (this Level) Allows(other Level) bool {
	return this >= other
}

func (this Level) String() string {
	return names[this]
}

func (this Level) FixedString() string {
	return fixedNames[this]
}
