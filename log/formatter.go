package log

import (
	"time"

	"github.com/sirupsen/logrus"
)

// DefaultTextFormatter returns a default formatter
func DefaultTextFormatter() logrus.Formatter {
	return &logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339, // "2006-01-02T15:04:05Z07:00"
	}
}

// LogstashFormatter represents a Logstash format.
// It has logrus.Formatter which formats the entry and logrus.Fields which
// are added to the JSON message if not given in the entry data.
//
// Note: use the `DefaultFormatter` function to set a default Logstash formatter.
type LogstashFormatter struct {
	logrus.Formatter
	logrus.Fields
}

var (
	logstashFields = logrus.Fields{
		"@version": "1",
		"type":     "log",
	}

	logstashFieldMap = logrus.FieldMap{
		logrus.FieldKeyTime: "@timestamp",
		logrus.FieldKeyMsg:  "message",
	}
)

// NewLogstashFormatter returns a default Logstash formatter:
// A JSON format with "@version" set to "1" (unless set differently in `fields`,
// "type" to "log" (unless set differently in `fields`),
// "@timestamp" to the log time and "message" to the log message.
//
// Note: to set a different configuration use the `LogstashFormatter` structure.
func NewLogstashFormatter(fields logrus.Fields) logrus.Formatter {
	for k, v := range logstashFields {
		if _, ok := fields[k]; !ok {
			fields[k] = v
		}
	}

	return LogstashFormatter{
		Formatter: &logrus.JSONFormatter{FieldMap: logstashFieldMap},
		Fields:    fields,
	}
}

func overwriteEntryFields(e *logrus.Entry, fields logrus.Fields) *logrus.Entry {
	return e.WithFields(fields)
}

// Format formats an entry to a Logstash format according to the given Formatter and Fields.
//
// Note: the given entry is copied and not changed during the formatting process.
func (f LogstashFormatter) Format(e *logrus.Entry) ([]byte, error) {
	ne := overwriteEntryFields(e, f.Fields)
	return f.Formatter.Format(ne)
}

func DefaultLogstashFormatter() logrus.Formatter {
	return NewLogstashFormatter(make(logrus.Fields))
}
