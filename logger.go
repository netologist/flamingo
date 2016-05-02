package vodka

import (
	"os"
	//"fmt"

	"github.com/Sirupsen/logrus"
	//"github.com/bshuster-repo/logrus-logstash-hook"
)

type Logger struct {
}

type LoggerFactory struct {
	fields map[string]interface{}
}

func NewLoggerFactory(fields map[string]interface{}) *LoggerFactory {
	return &LoggerFactory{fields}
}

func init() {

	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	logrus.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.WarnLevel)
}

func NewWithFields(fields map[string]interface{}) *logrus.Entry {
	log := logrus.New()
	/*  hook, err := logrus_logstash.NewHook("tcp", "172.17.0.2:9999")
	    if err != nil {
	            log.Fatal(err)
	    }
	    log.Hooks.Add(hook)
	*/
	ctx := log.WithFields(logrus.Fields{
		"method": "main",
	})

	return ctx
}
