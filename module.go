package logging

import "reflect"

type ModuleId string

func GetModule(value interface{}) ModuleId {
	return ModuleId(reflect.TypeOf(value).PkgPath())
}

func (this ModuleId) String() string {
	return string(this)
}
