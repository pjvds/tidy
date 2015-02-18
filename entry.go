package slog

import "time"

type Entry struct {
	Timestamp time.Time
	Module    ModuleId
	Level     Level

	Message string
	Fields  Fields
}
