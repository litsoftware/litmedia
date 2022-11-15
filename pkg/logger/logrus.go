package logger

import (
	"github.com/litsoftware/litmedia/internal/pkg/config"
	"github.com/litsoftware/litmedia/pkg/file"
	"github.com/litsoftware/litmedia/pkg/path"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var log *logrus.Logger
var logFile *os.File

func init() {
	logrus.RegisterExitHandler(closeLogFile)
	log = logrus.New()

	var level logrus.Level
	switch config.GetString("logging.level") {
	case "warning":
		level = logrus.WarnLevel
	case "error":
		level = logrus.ErrorLevel
	case "fatal":
		level = logrus.FatalLevel
	case "panic":
		level = logrus.PanicLevel
	case "debug":
		level = logrus.DebugLevel
	default:
		level = logrus.InfoLevel
	}

	log.SetOutput(os.Stdout)
	log.SetLevel(level)
	log.SetFormatter(&logrus.JSONFormatter{})

	//local()
	log.AddHook(lsfHook())
}

func local() {
	logPath := path.StoragePath() + "/logs/"
	logFile := logPath + "entrance.log"

	_ = file.MakeDirIfNotExist(logPath, 0755)

	lf, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.SetOutput(io.MultiWriter(os.Stdout, lf))
}

func lsfHook() *lfshook.LfsHook {
	logPath := path.StoragePath() + "/logs/"
	logName := logPath + "entrance.log"
	writer, _ := rotatelogs.New(
		logName+".%Y%m%d%H.log",
		rotatelogs.WithLinkName(logName),
		rotatelogs.WithRotationTime(time.Hour*time.Duration(24)),
		rotatelogs.WithRotationCount(50),
	)

	return lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{})
}

func closeLogFile() {
	if logFile != nil {
		_ = logFile.Close()
	}
}
