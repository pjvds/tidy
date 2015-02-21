package logging

import "time"

// Represents a single log entry.
type Entry struct {
	// The time this entry was created.
	Timestamp time.Time

	// The module this entry belongs to.
	Module ModuleId

	// The level of this entry, eq: DEBUG, WARN, FATAL.
	Level Level

	// The final message.
	Message string

	// The fields for this entry.
	Fields Fields
}
