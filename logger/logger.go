package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

// 供全局调用的不同日志级别
var (
	Error   *log.Logger
	Warning *log.Logger
	Info    *log.Logger
	Debug   *log.Logger
)

func init() {
	dir, _ := os.Getwd()
	systemLogPath := path.Join(dir, "logger", "system.log")
	systemLogFile, err := os.OpenFile(systemLogPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		panic(fmt.Sprintf("systemLogFile加载失败：%s", err))
	}

	// 设置日志输出位置和格式
	Error = log.New(io.MultiWriter(systemLogFile, os.Stderr), "[Error] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	Warning = log.New(io.MultiWriter(systemLogFile, os.Stderr), "[Warning] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	Info = log.New(io.MultiWriter(systemLogFile, os.Stderr), "[Info] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	Debug = log.New(io.MultiWriter(systemLogFile, os.Stderr), "[Debug] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
}
