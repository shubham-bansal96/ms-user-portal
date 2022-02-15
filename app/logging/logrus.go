package logging

import (
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/ms-user-portal/app/config"
	"github.com/sirupsen/logrus"
)

func Initialize(cfg *config.Configuration) {
	if level, err := logrus.ParseLevel(cfg.LogLevel); err == nil {
		logrus.SetLevel(level)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func LogForFunc() *logrus.Entry {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Could not get context info for logger!")
	}

	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]

	var hostname string
	hostName, err := os.Hostname()
	if err == nil && hostName != "" {
		hostname = hostName
	}
	return logrus.WithFields(logrus.Fields{
		"file":        filename,
		"function":    fn,
		"app-name":    config.Config.MSName,
		"environment": config.Config.Environment,
		"hostname":    hostname,
	})
}
