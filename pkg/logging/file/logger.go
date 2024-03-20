package logger

import (
	"log/slog"
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

const AccessLogFormat = "request_id=${id}, time=${time_rfc3339_nano}, method=${method}, host${host}, uri=${uri}, status=${status}, error=${error}, referer=${referer}, remote_ip=${remote_ip}, user_agent=${user_agent}, latency=${latency}, latency_human=${latency_human}, bytes_in=${bytes_in}, bytes_out=${bytes_out}\n"

var (
	log         *slog.Logger
	logDir      = "./log"
	filePath    = "/fiber.log"
	logSize     = 10   // MB
	logBucket   = 5    // 個数
	logAge      = 60   // days
	logCompress = true // 圧縮
)

func Initialize() {
	NewLogger()
}

func SetLogDir(dir string) {
	logDir = dir
}

func SetLogBucket(bucket int) {
	logBucket = bucket
}

func SetLogAge(age int) {
	logAge = age
}

func SetLogCompress(compress bool) {
	logCompress = compress
}

func NewLogger() {
	log = slog.New(slog.NewJSONHandler(&lumberjack.Logger{
		Filename:   filepath.Join(logDir, filePath),
		MaxSize:    logSize,
		MaxBackups: logBucket,
		MaxAge:     logAge,
		Compress:   logCompress,
	}, nil))
}

func Info(msg string, args ...interface{}) {
	log.Info(msg, args...)
}
func Warning(msg string, args ...interface{}) {
	log.Warn(msg, args...)
}
func Error(msg string, args ...interface{}) {
	log.Error(msg, args...)
}

// Fatal exit application
func Fatal(msg string, args ...interface{}) {
	log.Error(msg, args...)
	os.Exit(1)
}
func Debug(msg string, args ...interface{}) {
	log.Debug(msg, args...)
}
func With(args ...interface{}) {
	log = log.With(args...)
}
