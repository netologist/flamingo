package vodka

import (
	"github.com/Sirupsen/logrus"
)

func NewLoggerFactory(fields map[string]interface{}) *LoggerFactory {
	return &LoggerFactory{fields}
}

type LoggerFactory struct {
	fields map[string]interface{}
}

func (l *LoggerFactory) NewLogger(name string) Logger {
	return Logger{name, l.fields}
}

type Logger struct {
	name          string
	defaultFields map[string]interface{}
}

func (l *Logger) newEntry(fields map[string]interface{}) *logrus.Entry {
	lfields := logrus.Fields{}

	// logger Name
	lfields["name"] = l.name

	// default fields
	if len(l.defaultFields) > 0 {
		for key, val := range l.defaultFields {
			lfields[key] = val
		}
	}

	// custom fields
	if len(fields) > 0 {
		for key, val := range fields {
			lfields[key] = val
		}
	}

	return logrus.WithFields(lfields)
}

func (l *Logger) PanicWithFields(format string, fields logrus.Fields, args ...interface{}) {
	entry := l.newEntry(fields)
	entry.Panicf(format, args...)
}

func (l *Logger) Panic(format string, args ...interface{}) {
	l.PanicWithFields(format, nil, args...)
}

func (l *Logger) ErrorWithFields(format string, fields logrus.Fields, args ...interface{}) {
	entry := l.newEntry(fields)
	entry.Errorf(format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.ErrorWithFields(format, nil, args...)
}

func (l *Logger) WarnWithFields(format string, fields logrus.Fields, args ...interface{}) {
	entry := l.newEntry(fields)
	entry.Warnf(format, args...)
}

func (l *Logger) Warn(format string, args ...interface{}) {
	l.WarnWithFields(format, nil, args...)
}

func (l *Logger) InfoWithFields(format string, fields logrus.Fields, args ...interface{}) {
	entry := l.newEntry(fields)
	entry.Infof(format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.InfoWithFields(format, nil, args...)
}

func (l *Logger) DebugWithFields(format string, fields logrus.Fields, args ...interface{}) {
	entry := l.newEntry(fields)
	entry.Debugf(format, args...)
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.DebugWithFields(format, nil, args...)
}
