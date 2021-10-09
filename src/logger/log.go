package logger

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func Infor(msg string, kvs ...interface{}) {
	n := len(kvs)
	fields := make(log.Fields)
	for i := 0; i < n/2; i++ {
		fields[fmt.Sprintf("%v", kvs[2*i])] = kvs[2*i+1]
	}

	log.WithFields(fields).Info(msg)
}

func Error(err error, msg string, kvs ...interface{}) {
	n := len(kvs)
	fields := make(log.Fields)
	for i := 0; i < n/2; i++ {
		fields[fmt.Sprintf("%v", kvs[2*i])] = kvs[2*i+1]
	}
	fields["error"] = err

	log.WithFields(fields).Error(msg)
}
