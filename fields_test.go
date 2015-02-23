package tidy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldsInitialLen(t *testing.T) {
	var uninitialized Fields
	assert.Equal(t, 0, uninitialized.Len())

	assert.Equal(t, 1, Fields{
		"foo": "bar",
	}.Len())
}

func TestFieldsClone(t *testing.T) {
	fields := Fields{
		"foo": "bar",
		"baz": 42,
	}

	cloned := fields.Clone(0)
	assert.Equal(t, fields.Len(), cloned.Len())

	for key, value := range fields {
		assert.Equal(t, value, cloned[key])
	}
}

func TestFieldsCloneDoesNotEffectLen(t *testing.T) {
	fields := Fields{
		"foo": "bar",
		"baz": 42,
	}

	// specify a positive grow size to test
	// that this has no influence on the reported length
	cloned := fields.Clone(5)

	assert.Equal(t, fields.Len(), cloned.Len())
}
