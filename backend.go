package tidy

import (
	"os"
	"sync/atomic"
)

type mergeBackend []Backend

func leveledsToBackends(leveleds ...LeveledBackend) []Backend {
	backends := make([]Backend, len(leveleds))
	for index, leveled := range leveleds {
		backends[index] = leveled.Backend
	}
	return backends
}

func mergeBackends(backends ...Backend) Backend {
	return mergeBackend(backends)
}

func (this mergeBackend) Log(entry Entry) {
	for _, b := range this {
		b.Log(entry)
	}
}

func (this mergeBackend) Flush() error {
	for _, b := range this {
		b.Flush()
	}
	return nil
}

func NewRootBackend(level Level, backend Backend) *RootBackend {
	root := RootBackend{}
	root.level.Store(level)
	root.backend.Store(backend)

	return &root
}

type wrappedBackend struct {
	Backend
}

type RootBackend struct {
	level   atomic.Value
	backend atomic.Value
}

func (this *RootBackend) ChangeLevel(level Level) {
	this.level.Store(level)
}

func (this *RootBackend) ChangeBackend(backend Backend) {
	this.backend.Store(wrappedBackend{backend})
}

func (this *RootBackend) IsEnabledFor(level Level, module Module) bool {
	// TODO: respect module
	return this.level.Load().(Level).Allows(level)
}
func (this *RootBackend) Log(entry Entry) {
	if this.IsEnabledFor(entry.Level, entry.Module) {
		wrapper := this.backend.Load()
		if wrapper == nil {
			// no backend was stored
			return
		}
		backend := wrapper.(wrappedBackend).Backend
		backend.Log(entry)
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
