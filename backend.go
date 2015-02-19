package logging

import "os"

type ColoredConsoleBackend struct {
	formatter ColoredTextFormatter
	Level     Level
}

func (this *ColoredConsoleBackend) IsEnabledFor(level Level, module ModuleId) bool {
	return this.Level >= level
}

func (this *ColoredConsoleBackend) Log(entry Entry) {
	this.formatter.FormatTo(os.Stderr, entry)
}

func (this *ColoredConsoleBackend) Flush() error {
	return os.Stderr.Sync()
}

type Backend interface {
	IsEnabledFor(level Level, module ModuleId) bool
	Log(entry Entry)
	Flush() error
}
