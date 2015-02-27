package tidy

import (
	"errors"
	"reflect"
	"runtime"
	"strings"
)

type Module string

func GetModuleFromValue(value interface{}) Module {
	return Module(reflect.TypeOf(value).PkgPath())
}

func GetModuleFromCaller(depth int) Module {
	pc, _, _, ok := runtime.Caller(1 + depth)

	if !ok {
		panic(errors.New("failed to get caller from runtime"))
	}

	function := runtime.FuncForPC(pc)

	if function == nil {
		panic(errors.New("failed to get function from program counter"))
	}

	// The function name is the complete package path and function name without signature seperated by a dot.
	// e.q.: github.com/pjvds/tidy.GetLogger
	name := function.Name()

	lastSlashIndex := strings.LastIndex(name, "/")
	lastDotIndex := strings.LastIndex(name, ".")

	return Module(name[lastSlashIndex+1 : lastDotIndex])
}

func (this Module) String() string {
	return string(this)
}
