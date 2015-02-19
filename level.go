package logging

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

func (this Level) String() string {
	return names[this]
}

func (this Level) FixedString() string {
	return fixedNames[this]
}
