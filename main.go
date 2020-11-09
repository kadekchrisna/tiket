package main

import (
	loggers "tiket.vip/src/infrastructures/logger"
	"tiket.vip/src/infrastructures/server"
)

func main() {
	loggers.Init()
	// logrus.WithFields(logrus.Fields{
	// 	"Animal": "Logrus",
	// }).Info("A logrus appears")

	server.Serve()
}
