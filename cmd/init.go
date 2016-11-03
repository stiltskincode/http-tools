package cmd

import (
	"github.com/Sirupsen/logrus"
	"os"
)



func init() {
	log.Out = os.Stdout
	log.Formatter = new(logrus.JSONFormatter)
}

func Logger(newlog *logrus.Logger) {
	log = newlog
}
