package main

import (
	. "github.com/mundipagg/logrus_seq_hook"
	"github.com/sirupsen/logrus"
)

func main() {
	seqHook, _ := NewSeqHook("http://localhost:5341", "4jZzTybZ9bUHtJiPdh6")
	logrus.AddHook(seqHook)
	logrus.WithFields(logrus.Fields{"key": "value"}).Info("Logging on SEQ")
	logrus.WithFields(logrus.Fields{"key": "value"}).Warn("Logging on SEQ")
	logrus.WithFields(logrus.Fields{"key": "value"}).Fatal("Logging on SEQ")

}
