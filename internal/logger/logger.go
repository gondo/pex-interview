package logger

import "fmt"

var debugEnabled = false

func EnableDebug() {
	debugEnabled = true
}

func Log(msg string) {
	if debugEnabled {
		fmt.Println(msg)
	}
}
