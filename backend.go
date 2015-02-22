package logging

import "os"

type LeveledBackend struct {
	Level Level
}

func (this LeveledBackend) IsEnabledFor(level Level, module ModuleId) bool {
	return this.Level >= level
}

type ColoredConsoleBackend struct {
	LeveledBackend

	formatter ColoredTextFormatter
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
