package tidy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Stringify(object interface{}) string {
	return fmt.Sprintf("%+v", object)
}

func TestStringify(t *testing.T) {
	result := Stringify(struct {
		Foo string
		Bar int
	}{
		Foo: "baz",
		Bar: 42,
	})

	assert.Equal(t, result, "foo")
}
