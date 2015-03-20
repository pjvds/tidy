package tidy

type Verbosity struct {
	enabled bool
	level   Level
	logger  *Logger
	fields  Fields
}

func newVerbosity(level Level, logger *Logger) Verbosity {
	return Verbosity{
		enabled: logger.IsEnabled(level),
		level:   level,
		logger:  logger,
	}
}

func (this Verbosity) With(field string, value interface{}) Verbosity {
	if this.enabled {
		this.fields[field] = value
	}
	return this
}

func (this Verbosity) Withs(fields Fields) Verbosity {
	if this.enabled {
		// no need to be immutable here, Verbosity types are there
		// to only create a single log line.
		for field, value := range fields {
			this.fields[field] = value
		}
	}
	return this
}

func (this Verbosity) Write(message string) Verbosity {
	if this.enabled {
		this.logger.Withs(this.fields).log(this.level, message)
	}
}
