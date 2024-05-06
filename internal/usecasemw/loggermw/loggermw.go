package loggermw

import (
	"context"

	"github.com/sirupsen/logrus"
)

func log(_ context.Context, funcName string, req any, res any, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"funcName": funcName,
		"req":      req,
		"res":      res,
		"err":      err,
	})

	level := logrus.InfoLevel
	if err != nil {
		level = logrus.ErrorLevel
	}

	logger.Log(level, level)
}
