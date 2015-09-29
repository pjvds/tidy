package tidy

import (
	"os"
	"sync/atomic"
)

func NewRootBackend(level Level, backend Backend) *RootBackend {
	root := RootBackend{}
	root.level.Store(level)
	root.backend.Store(backend)

	return &root
}

type RootBackend struct {
	level   atomic.Value
	backend atomic.Value
}

func (this *RootBackend) ChangeLevel(level Level) {
	this.level.Store(level)
}

func (this *RootBackend) ChangeBackend(backend Backend) {
	this.backend.Store(backend)
}

func (this *RootBackend) IsEnabledFor(level Level, module Module) bool {
	// TODO: respect module
	return this.level.Load().(Level).Allows(level)
}
func (this *RootBackend) Log(entry Entry) {
	if this.IsEnabledFor(entry.Level, entry.Module) {
		this.backend.Load().(Backend).Log(entry)
	}
}

func (this *RootBackend) Flush() error {
	return this.backend.Load().(Backend).Flush()
}

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
