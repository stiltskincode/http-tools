package cmd

import (
	"github.com/Sirupsen/logrus"
	"os"
	"time"
	"math/rand"
)

func init() {
	rand.Seed(time.Now().Unix())
	log.Out = os.Stdout
	log.Formatter = new(logrus.JSONFormatter)
}

func Logger(newlog *logrus.Logger) {
	log = newlog
}
