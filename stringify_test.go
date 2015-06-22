package tidy_test

import (
	"testing"

	"github.com/pjvds/tidy"
	"github.com/stretchr/testify/assert"
)

func TestStringify(t *testing.T) {
	assert.Equal(t, "{Foo:baz Bar:42}", tidy.Stringify(struct {
		Foo string
		Bar int
	}{
		Foo: "baz",
		Bar: 42,
	}))
}

func TestStringifyPrintsPrivates(t *testing.T) {
	assert.Equal(t, "{Foo:baz Bar:42 private:private}", tidy.Stringify(struct {
		Foo     string
		Bar     int
		private string
	}{
		Foo:     "baz",
		Bar:     42,
		private: "private",
	}))
}
