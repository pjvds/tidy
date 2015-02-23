package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocation(t *testing.T) {
	location := GetLocation(0)

	assert.Equal(t, location.String(), "location_test.go:10")
}

func BenchmarkGetLocation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetLocation(0)
	}
}
