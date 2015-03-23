package tidy

type verbosity struct {
	enabled bool
	level   Level
	logger  Logger
	fields  Fields
}

func (this verbosity) With(field string, value interface{}) verbosity {
	if this.enabled {
		this.fields[field] = value
	}
	return this
}

func (this verbosity) Withs(fields Fields) verbosity {
	if this.enabled {
		// no need to be immutable here, verbosity types are there
		// to only create a single log line.
		for field, value := range fields {
			this.fields[field] = value
		}
	}
	return this
}

func (this verbosity) Write(message string) verbosity {
	if this.enabled {
		this.logger.Withs(this.fields).log(this.level, message)
	}

	return this
}
