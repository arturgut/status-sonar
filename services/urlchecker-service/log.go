package main

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func initLogging() {

	log.SetLevel(log.DebugLevel) // Only log the warning severity or above.

	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05", // the "time" field configuratiom
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
	}
	logrus.SetFormatter(formatter)

	log.Info("Starting URL checker. Version: 0.1")
}
