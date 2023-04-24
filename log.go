package utils

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"
	"time"
)

type LoggerManage struct {
	logger *logrus.Logger
}

var Logger LoggerManage

func init() {
	logrusLogger := logrus.New()

	logFile := "log.log"
	// 每隔 4MB 分割日志文件
	rotateInterval := 4 * 1024 * 1024

	// 设置日志输出格式
	logrusLogger.Formatter = &logrus.TextFormatter{}

	writer, err := rotatelogs.New(
		logFile+".%Y%m%d%H%M%S",
		rotatelogs.WithLinkName(logFile),
		rotatelogs.WithRotationSize(int64(rotateInterval)),
		rotatelogs.WithMaxAge(time.Hour*24*7),
	)
	if err != nil {
		fmt.Printf("Failed to create log rotate: %v", err)
		os.Exit(-1)
	}

	logWriter := io.MultiWriter(os.Stdout, writer)

	logrusLogger.SetOutput(logWriter)
	Logger.logger = logrusLogger
}

func (l LoggerManage) Warnln(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.logger.Warnln(file, line, args)
}

func (l LoggerManage) Infoln(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.logger.Infoln(file, line, args)
}
func (l LoggerManage) Errorln(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.logger.Errorln(file, line, args)
}
func (l LoggerManage) Debugln(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.logger.Debugln(file, line, args)
}
func (l LoggerManage) Errorf(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.logger.Errorf(file, line, args)
}
