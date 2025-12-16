package logger

import "log"

func _log(level string, msg ...any) {
	args := append(msg, level)
	log.Println(args...)
}

func Info(msg ...any) {
	_log("[INFO]", msg...)
}

func Warn(msg ...any) {
	_log("[INFO]", msg...)
}

func Err(msg ...any) {
	_log("[INFO]", msg...)
}
