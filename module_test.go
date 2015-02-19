package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyStructure struct{}

func TestGetModule(t *testing.T) {
	module := GetModule(MyStructure{})

	assert.Equal(t, module.String(), "github.com/pjvds/logging")
}
