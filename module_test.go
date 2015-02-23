package tidy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var moduleAtInit = GetModuleFromCaller(0)

type MyStructure struct{}

func TestGetModule(t *testing.T) {
	module := GetModule(MyStructure{})

	assert.Equal(t, module.String(), "github.com/pjvds/tidy")
}

func TestGetModuleFromCaller(t *testing.T) {
	module := GetModuleFromCaller(0)

	assert.Equal(t, module.String(), "tidy")
}

func TestGetModuleFromCallerAtInit(t *testing.T) {
	assert.Equal(t, moduleAtInit.String(), "tidy")
}
