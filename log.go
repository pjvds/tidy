package logging

import (
	"fmt"
	"time"
)

type Logger struct {
	module Module
	fields Fields

	backend Backend
}

var defaulBackend Backend

func init() {
	defaulBackend = &ColoredConsoleBackend{
		formatter: ColoredTextFormatter{},
	}
}

func CreateOrGetLogger(module string) *Logger {
	return &Logger{
		module:  Module(module),
		backend: defaulBackend,
	}
}

func (this *Logger) WithField(key string, value interface{}) *Logger {
	clone := make(Fields, len(this.fields)+1)

	for existingKey, existingValue := range this.fields {
		clone[existingKey] = existingValue
	}

	clone[key] = value

	return &Logger{
		module:  this.module,
		fields:  clone,
		backend: this.backend,
	}
}

func (this *Logger) WithFields(fields Fields) *Logger {
	return &Logger{
		module:  this.module,
		fields:  this.fields.Join(fields),
		backend: this.backend,
	}
}

func (this *Logger) IsDebug() bool {
	return this.backend.IsEnabledFor(DEBUG, this.module)
}
func (this *Logger) IsInfo() bool {
	return this.backend.IsEnabledFor(INFO, this.module)
}
func (this *Logger) IsWarn() bool {
	return this.backend.IsEnabledFor(WARN, this.module)
}
func (this *Logger) IsError() bool {
	return this.backend.IsEnabledFor(ERROR, this.module)
}

func (this *Logger) log(level Level, msg string) {
	this.backend.Log(Entry{
		Timestamp: time.Now(),
		Module:    this.module,
		Level:     level,
		Message:   msg,
		Fields:    this.fields,
	})
}

func (this *Logger) Debug(msg string) {
	this.log(DEBUG, msg)
}
func (this *Logger) Debugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	this.Debug(msg)
}
func (this *Logger) Info(msg string) {
	this.log(INFO, msg)
}
func (this *Logger) Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	this.Info(msg)
}
func (this *Logger) Warn(msg string) {
	this.log(WARN, msg)
}
func (this *Logger) Warnf(format string, args ...interface{}) {
	this.log(WARN, fmt.Sprintf(format, args...))
}
func (this *Logger) Error(msg string) {
	this.log(ERROR, msg)
}
func (this *Logger) Errorf(format string, args ...interface{}) {
	this.log(ERROR, fmt.Sprintf(format, args...))
}
func (this *Logger) Fatal(msg string) {
	this.log(FATAL, msg)
}
func (this *Logger) Fatalf(format string, args ...interface{}) {
	this.log(FATAL, fmt.Sprintf(format, args...))
}
