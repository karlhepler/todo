package service

import "github.com/oldtimeguitarguy/todo/app/service/driver"

type Logger struct {
	driver.Logger
}

func (srv Logger) LogError(err error) {
	srv.Logger.Printf("[ERROR] %s", err.Error())
}
