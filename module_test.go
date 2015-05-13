package tidy_test

import (
	"testing"

	"github.com/pjvds/tidy"
	"github.com/stretchr/testify/assert"
)

var moduleAtInit = tidy.GetModuleFromCaller(0)

type MyStructure struct{}

func TestGetModuleFromValue(t *testing.T) {
	module := tidy.GetModuleFromValue(MyStructure{})

	assert.Equal(t, module.String(), "tidy_test")
}

func TestGetModuleFromCaller(t *testing.T) {
	module := tidy.GetModuleFromCaller(0)

	assert.Equal(t, module.String(), "tidy_test")
}

func TestGetModuleFromCallerAtInit(t *testing.T) {
	assert.Equal(t, moduleAtInit.String(), "tidy_test")
}

func TestGetModuleFromInlineFunc(t *testing.T) {
	assert.Equal(t, "tidy_test", func() tidy.Module {
		return tidy.GetModuleFromCaller(0)
	}().String())
}

type foo struct{}

func (this foo) GetModule() tidy.Module {
	return tidy.GetModuleFromCaller(0)
}

func TestGetModuleFromAttachedFunction(t *testing.T) {
	assert.Equal(t, "tidy_test", foo{}.GetModule().String())
}
