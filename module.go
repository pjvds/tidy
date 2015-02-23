package logging

import (
	"errors"
	"reflect"
	"runtime"
	"strings"
)

type ModuleId string

func GetModule(value interface{}) ModuleId {
	return ModuleId(reflect.TypeOf(value).PkgPath())
}

func GetModuleFromCaller(depth int) ModuleId {
	pc, _, _, ok := runtime.Caller(1 + depth)

	if !ok {
		panic(errors.New("failed to get caller from runtime"))
	}

	function := runtime.FuncForPC(pc)

	if function == nil {
		panic(errors.New("failed to get function from program counter"))
	}

	name := function.Name()
	lastSlashIndex := strings.LastIndex(name, "/")
	lastDotIndex := strings.LastIndex(name, ".")

	return ModuleId(name[lastSlashIndex+1 : lastDotIndex])
}

func (this ModuleId) String() string {
	return string(this)
}
