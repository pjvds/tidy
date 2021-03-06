package tidy

import (
	"errors"
	"os"
)

func Configure() config {
	return config{}
}

type config struct {
	backends []LeveledBackend
}

type toBuilder struct {
	level Level
	cfg   config
}

func (this config) LogFromLevelSpecifiedByEnvironment() toBuilder {
	level := ParseLevel(os.Getenv("LOG"))

	return toBuilder{
		level: level,
		cfg:   this,
	}
}

func (this toBuilder) To(backend BackendBuilder) config {
	cfg := this.cfg
	cfg.backends = append(cfg.backends, NewLeveledBackend(this.level, backend.Build()))
	return cfg
}

func (this config) LogFromLevel(level Level) toBuilder {
	return toBuilder{
		level: level,
		cfg:   this,
	}
}

func (this config) BuildDefault() error {
	if len(this.backends) == 0 {
		return errors.New("no backend found in config, forgot Configure().To() call?")
	}

	lowestLevel := this.backends[0].Level

	for _, backend := range this.backends {
		if backend.Level.Allows(lowestLevel) {
			lowestLevel = backend.Level
		}
	}
	merged := mergeBackends(leveledsToBackends(this.backends...)...)

	defaulBackend.ChangeLevel(lowestLevel)
	defaulBackend.ChangeBackend(merged)
	return nil
}

func (this config) MustBuildDefault() {
	if err := this.BuildDefault(); err != nil {
		panic(err)
	}
}

func (this config) Build() (Logger, error) {
	if len(this.backends) == 0 {
		return Logger{}, errors.New("no backend found in config, forgot Configure().To() call?")
	}

	return NewLogger(GetModuleFromCaller(2), this.backends[0]), nil
}

func (this config) MustBuild() Logger {
	logger, err := this.Build()

	if err != nil {
		panic(err)
	}

	return logger
}

type BackendBuilder interface {
	Build() Backend
}
