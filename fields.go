package logging

type Fields map[string]interface{}

func (this Fields) Len() int {
	return len(this)
}

func (this Fields) Any() bool {
	return len(this) > 0
}

func (this Fields) Clone(grow int) Fields {
	clone := make(Fields, this.Len()+grow)

	for key, value := range this {
		clone[key] = value
	}

	return clone
}

func (this Fields) Join(fields Fields) Fields {
	joined := this.Clone(fields.Len())

	for key, value := range fields {
		joined[key] = value
	}

	return joined
}
