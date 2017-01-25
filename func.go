package testlib

import (
	"reflect"
	"runtime"
	"strings"
)

func FullFunctionName(i interface{}) string {
	full := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return full
}

func FunctionName(i interface{}) string {
	full := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	name := full[strings.LastIndex(full, ".")+1:]
	return name
}
