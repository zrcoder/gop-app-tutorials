package log

import (
	"fmt"
	"path/filepath"
	"runtime"
)

var (
	IsDev = true
)

func Println(v ...interface{}) {
	if !IsDev {
		return
	}

	printTrace()
	fmt.Println(v...)
}

func Printf(format string, v ...interface{}) {
	if !IsDev {
		return
	}
	printTrace()
	fmt.Printf(format, v...)
}

func printTrace() {
	_, file, line, _ := runtime.Caller(2)
	_, file = filepath.Split(file)
	fmt.Printf("[%s:%d] ", file, line)
}
