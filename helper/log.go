package helper

import (
	"log"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
	"github.com/titoyudha/go_ecommerce_api/config"
)

var once sync.Once

var Log *logrus.Logger

func init() {
	once.Do(func() {
		Log = newLogger()
	})
}

func newLogger() *logrus.Logger {
	Log := logrus.New()
	log.Println("Setup Logger")
	if config.Get().LogPath != "" {
		err := os.Mkdir(config.Get().LogPath, 0755)
		if err != nil {
			log.Println("Failed to create log path")
		}
		log.Println("Success Creating Log Path")
	}

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   config.Get().LogPath,
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     30, //in days
		Compress:   false,
		Level:      0,
		Formatter: &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02T15:04:05.999999999Z07:00",
			ForceColors:     true,
		},
	})
	if err != nil {
		log.Fatalf("Failed to initialize file rotation hook: %v", err)
	}
	Log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	Log.SetReportCaller(true)
	Log.AddHook(rotateFileHook)
	return Log
}
