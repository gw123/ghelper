package grpc

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"time"
)

func initGRPCLogrus() (*logrus.Entry, []grpc_logrus.Option) {
	// https://godoc.org/github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())

	logrusEntry.Logger.Formatter = &logrus.JSONFormatter{}

	// Shared options for the logger, with a custom gRPC grpc_code to log level function.
	option := []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}

	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	return logrusEntry, option
}
