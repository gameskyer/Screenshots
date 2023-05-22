package util

import (
	"Screenshots/db"
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//日志自定义格式
type LogFormatter struct{}

var Logger = logrus.New()

//格式详情
func (s *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05.000")
	var file string
	var len int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		len = entry.Caller.Line
	}
	//fmt.Println(entry.Data)
	msg := fmt.Sprintf("%s [%s:%d][%s] %s\n", timestamp, file, len, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}
func init() {
	path := "go.log"
	/* 日志轮转相关函数
	`WithLinkName` 为最新的日志建立软连接
	`WithRotationTime` 设置日志分割的时间，隔多久分割一次
	WithMaxAge 和 WithRotationCount二者只能设置一个
	 `WithMaxAge` 设置文件清理前的最长保存时间
	 `WithRotationCount` 设置文件清理前最多保存的个数
	*/
	// 下面配置日志每隔一天轮转一个新文件，保留最近10天的日志文件，多余的自动清理掉。
	logsPath := "templates/static/logs/"
	if !db.Exists(logsPath) {
		os.MkdirAll(logsPath, os.ModePerm)
	}
	writer, _ := rotatelogs.New(
		logsPath+path+".%Y%m%d",
		rotatelogs.WithLinkName(logsPath+path),
		rotatelogs.WithRotationCount(10),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.FatalLevel: writer,
		logrus.DebugLevel: writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.PanicLevel: writer,
	}
	// 显示行号等信息
	Logger.SetReportCaller(true)
	Logger.SetFormatter(new(LogFormatter))
	// 设置日志级别为InfoLevel以及以上
	Logger.SetLevel(logrus.DebugLevel)
	lfHook := lfshook.NewHook(writeMap, new(LogFormatter))
	Logger.AddHook(lfHook)
}
