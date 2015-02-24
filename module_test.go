package tidy_test

import (
	"testing"

	"github.com/pjvds/tidy"
	"github.com/stretchr/testify/assert"
)

var moduleAtInit = tidy.GetModuleFromCaller(0)

type MyStructure struct{}

func TestGetModule(t *testing.T) {
	module := tidy.GetModule(MyStructure{})

	assert.Equal(t, module.String(), "github.com/pjvds/tidy_test")
}

func TestGetModuleFromCaller(t *testing.T) {
	module := tidy.GetModuleFromCaller(0)

	assert.Equal(t, module.String(), "tidy_test")
}

func TestGetModuleFromCallerAtInit(t *testing.T) {
	assert.Equal(t, moduleAtInit.String(), "tidy_test")
}
