package log

import (
	"io/ioutil"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
)

const (
	MaxSize    = 100
	MaxBackups = 7
	MaxAge     = 7
)

func FileLogger(path string) (*logrus.Logger, error) {
	logger := logrus.New()
	logger.SetOutput(ioutil.Discard)

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   path,
		MaxSize:    MaxSize,
		MaxBackups: MaxBackups,
		MaxAge:     MaxAge,
		Level:      logrus.InfoLevel,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat:  time.RFC3339,
			DisableTimestamp: false,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg:  "message",
				logrus.FieldKeyTime: "timestamp",
			},
		},
	})
	if err != nil {
		return nil, err
	}

	logger.AddHook(rotateFileHook)

	return logger, nil
}
