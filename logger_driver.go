package main

import "log"

type MyLogger struct {
	Log *log.Logger
}

func (l MyLogger) LogError(err error) {
	l.Log.Printf("ERROR: %s", err.Error())
}
