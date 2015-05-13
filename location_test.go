package tidy_test

import (
	"testing"

	"github.com/pjvds/tidy"
	"github.com/stretchr/testify/assert"
)

func TestGetLocation(t *testing.T) {
	location := tidy.GetLocation(0)

	assert.Equal(t, location.String(), "location_test.go:11")
}

func TestLocationIsEmpty(t *testing.T) {
	notEmpty := tidy.Location("location_test.go:11")
	empty := tidy.Location("")

	assert.False(t, notEmpty.IsEmpty())
	assert.True(t, empty.IsEmpty())
}

func BenchmarkGetLocation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tidy.GetLocation(0)
	}
}
