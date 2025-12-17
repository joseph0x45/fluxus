package logger

import "log"

func _log(level string, msg ...any) {
	args := append([]any{level}, msg...)
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
