package tidy

import (
	"fmt"
	"os"
	"time"
)

// Immutable logger context.
type Logger struct {
	module Module
	fields Fields

	backend LeveledBackend
}

var defaulBackend Backend

func init() {
	defaulBackend = &ColoredConsoleBackend{
		LeveledBackend: LeveledBackend{
			Level: DEBUG,
		},
		formatter: ColoredTextFormatter{},
	}
}

func GetLogger() Logger {
	module := GetModuleFromCaller(1)
	logger := NewLogger(module, defaulBackend)

	return logger
}

func CreateOrGetLogger(module string) Logger {
	return Logger{
		module: NewModule(module),

		// TODO: make configurable
		backend: NewLeveledBackend(DEBUG, defaulBackend),
	}
}

func NewLogger(module Module, backend Backend) Logger {
	return Logger{
		module: module,

		// TODO: make configurable
		backend: NewLeveledBackend(DEBUG, backend),
	}
}

// Obsolete: use With instead.
func (this Logger) WithField(key string, value interface{}) Logger {
	return this.With(key, value)
}

// Obsolete: use Withs instead.
func (this Logger) WithFields(fields Fields) Logger {
	return this.Withs(fields)
}

// With return a copy of the current Logger with the specified field set
// to the specified value.
func (this Logger) With(key string, value interface{}) Logger {
	clone := make(Fields, len(this.fields)+1)

	for existingKey, existingValue := range this.fields {
		clone[existingKey] = existingValue
	}

	clone[key] = value

	return Logger{
		module:  this.module,
		fields:  clone,
		backend: this.backend,
	}
}

// Withs returns a copy of the current Logger with the additional specified fields.
func (this Logger) Withs(fields Fields) Logger {
	return Logger{
		module:  this.module,
		fields:  this.fields.Join(fields),
		backend: this.backend,
	}
}

func (this Logger) IsEnabled(level Level) bool {
	return this.backend.IsEnabledFor(level, this.module)
}

func (this Logger) IsDebug() bool {
	return this.backend.IsEnabledFor(DEBUG, this.module)
}
func (this Logger) IsInfo() bool {
	return this.backend.IsEnabledFor(INFO, this.module)
}
func (this Logger) IsWarn() bool {
	return this.backend.IsEnabledFor(WARN, this.module)
}
func (this Logger) IsError() bool {
	return this.backend.IsEnabledFor(ERROR, this.module)
}

func (this Logger) log(level Level, msg string) {
	this.backend.Log(Entry{
		Timestamp: time.Now(),
		Module:    this.module,
		Level:     level,
		Message:   msg,
		Fields:    this.fields,
	})
}

func (this Logger) Debug(msg string) {
	this.log(DEBUG, msg)
}
func (this Logger) Debugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	this.Debug(msg)
}
func (this Logger) Info(msg string) {
	this.log(INFO, msg)
}
func (this Logger) Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	this.Info(msg)
}
func (this Logger) Warn(msg string) {
	this.log(WARN, msg)
}
func (this Logger) Warnf(format string, args ...interface{}) {
	this.log(WARN, fmt.Sprintf(format, args...))
}
func (this Logger) Error(msg string) {
	this.log(ERROR, msg)
}
func (this Logger) Errorf(format string, args ...interface{}) {
	this.log(ERROR, fmt.Sprintf(format, args...))
}

// Fatal logs a FATAL message and then calls os.Exit(255). This causes the current program to
// exit with the 255 status code. Conventionally, code zero indicates success, non-zero an error.
// The program terminates immediately; deferred functions are not run.
func (this Logger) Fatal(msg string) {
	this.log(FATAL, msg)
	os.Exit(255)
}

// Fatalf formats the messages and logs a FATAL message and then calls os.Exit(255). This causes the current program to
// exit with the 255 status code. Conventionally, code zero indicates success, non-zero an error.
// The program terminates immediately; deferred functions are not run.
func (this Logger) Fatalf(format string, args ...interface{}) {
	this.log(FATAL, fmt.Sprintf(format, args...))
	os.Exit(255)
}

func (this Logger) Panic(err error) {
	this.log(ERROR, err.Error())
	panic(err)
}

func (this Logger) Panicf(format string, args ...interface{}) {
	err := fmt.Errorf(format, args...)
	this.log(ERROR, err.Error())
	panic(err)
}

func (this Logger) V(level Level) verbosity {
	return verbosity{
		enabled: this.IsEnabled(level),
		level:   level,
		logger:  this,
	}
}
