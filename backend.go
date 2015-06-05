package tidy

import "os"

type LeveledBackend struct {
	Level Level

	Backend
}

func NewLeveledBackend(level Level, backend Backend) LeveledBackend {
	return LeveledBackend{
		Level:   level,
		Backend: backend,
	}
}

func (this LeveledBackend) IsEnabledFor(level Level, module Module) bool {
	// TODO: respect module
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
	Log(entry Entry)
	Flush() error
}
