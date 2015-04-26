package tidy

// TODO:
// this file is VERY EXPERIMENTAL!
// it is just a tryout to find a good API to prevent unnecessary argument evaluation.

type verbosity struct {
	enabled bool
	level   Level
	logger  Logger
	fields  Fields
}

func (this verbosity) With(field string, value interface{}) verbosity {
	if this.enabled {
		if this.fields == nil {
			this.fields = make(Fields, 3)
		}
		this.fields[field] = value
	}
	return this
}

func (this verbosity) Withs(fields Fields) verbosity {
	if this.enabled {
		if this.fields == nil {
			// when we don't have fields yet,
			// just use the provided.
			this.fields = fields
		} else {
			// no need to be immutable here, verbosity types are there
			// to only create a single log line.
			for field, value := range fields {
				this.fields[field] = value
			}
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
