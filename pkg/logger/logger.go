package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

const SERVER = "server"
const DEFAULT = "default"

var Logger *logrus.Logger
var Date string

func GetLogger() *logrus.Logger {
	now := time.Now()
	nowDate := now.Format("2006-01-02")

	if Date == nowDate && Logger != nil {
		return Logger
	}

	// 開資料夾
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}

	//日志文件
	logFileName := nowDate + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		DisableHTMLEscape: true,
		TimestampFormat:   "2006-01-02 15:04:05",
	})

	// 設置單例
	Logger = logger
	Date = nowDate

	return logger
}

func Info(log interface{}, category string) {
	logger := GetLogger()
	logger.WithFields(logrus.Fields{
		"category": category,
	}).Info(log)
}

func Error(log interface{}, category string) {
	logger := GetLogger()
	logger.WithFields(logrus.Fields{
		"category": category,
	}).Error(log)
}
