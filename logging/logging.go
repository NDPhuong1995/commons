package logging

import (
	"fmt"
	"github.com/grasshopper/commons/filesystem"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Level int

var (
	F *os.File
	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func SetUp(dir *string) {
	if dir == nil {
		logger = log.New(os.Stderr, DefaultPrefix, log.LstdFlags)
	} else {
		fileName := fmt.Sprintf("%s%s.%s", "logging", time.Now().Format("20060102"), "log")

		F, err := filesystem.MustOpenFile(fileName, *dir + fileName)
		if err != nil {
			log.Fatalf("logging.Setup err: %v", err)
		}
		logger = log.New(F, DefaultPrefix, log.LstdFlags)
	}
}

func Debug(msg string, v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(msg, v)
}

func Info(msg string, v ...interface{}) {
	setPrefix(INFO)
	logger.Println(msg, v)
}

func Warn(msg string, v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(msg, v)
}

func Error(msg string, v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(msg, v)
}

func Fatal(msg string, v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(msg, v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}