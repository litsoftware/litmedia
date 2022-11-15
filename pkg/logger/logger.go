package logger

import (
	"github.com/sirupsen/logrus"
)

type Level = logrus.Level
type Logger struct {}

func (*Logger) Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func (*Logger) Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func (*Logger) Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (*Logger) Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func (*Logger) Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func (*Logger) Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
	log.Exit(1)
}

func (*Logger) Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func (*Logger) Debug(args ...interface{}) {
	log.Debug(args...)
}

func (*Logger) Info(args ...interface{}) {
	log.Info(args...)
}

func (*Logger) Print(args ...interface{}) {
	log.Print(args...)
}

func (*Logger) Warn(args ...interface{}) {
	log.Warn(args...)
}

func (*Logger) Error(args ...interface{}) {
	log.Error(args...)
}

func (*Logger) Fatal(args ...interface{}) {
	log.Fatal(args...)
	log.Exit(1)
}

func (*Logger) Panic(args ...interface{}) {
	log.Panic(args...)
}

func GetLogger() *Logger {
	return new(Logger)
}