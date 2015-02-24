package tidy

// Fields is a simple bag of key and value pairs that can
// be added to a log. Mostly used as an argument to the
// `Logger.WithFields` method. Fields can be created like a map:
//
//     fields := tidy.Fields{
//         "username": "foobar",
//         "count": 42,
//     }
type Fields map[string]interface{}

// Len returns the number of fields in the instance.
func (this Fields) Len() int {
	return len(this)
}

// Any returns true when there is any field in this instance; otherwise false.
func (this Fields) Any() bool {
	return len(this) > 0
}

// Clone creates a new Fields instance that holds the same values
// as this instance with an increased capacity this instance plus
// the specified grow size.
func (this Fields) Clone(grow int) Fields {
	clone := make(Fields, this.Len()+grow)

	for key, value := range this {
		clone[key] = value
	}

	return clone
}

// Join creates a new Fields instance with the content of this instance and
// the specified one.
func (this Fields) Join(fields Fields) Fields {
	joined := this.Clone(fields.Len())

	for key, value := range fields {
		joined[key] = value
	}

	return joined
}
