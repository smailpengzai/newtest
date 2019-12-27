package main

import (
	"fmt"
	"net/http"
)

type Logger struct{}

func (this *Logger) Debug() {
	if this == nil {
		panic("fuck")
	}
	fmt.Println("hello world")
}

var __logger *Logger

func AppLog() *Logger {
	return __logger
}

func InitConf() {
	__logger = &Logger{}
}

// ------内上内容应该是写在log包中，这里是方便演示------

var logger = AppLog()

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	AppLog().Debug() // ok
	logger.Debug()   // panic
}

func main() {
	InitConf()
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":8088", nil)
}
