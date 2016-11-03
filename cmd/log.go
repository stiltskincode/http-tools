package cmd

import (
	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

type BenchmarkLog struct {
	Fields logrus.Fields
	Message string
}

func logDispatch() {
	go func() {
		for {
			select {
			case b := <-LogChannel:
				log.WithFields(b.Fields).Info(b.Message)
			}
		}
	}()
}

