package logger

import (
	"github.com/go-playground/log"
	"github.com/go-playground/log/handlers/console"
)

func Setup() {
	cLog := console.New(true)
	log.AddHandler(cLog, log.AllLevels...)
}
