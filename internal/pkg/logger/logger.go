package logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.Logger

func InitLogger() {

	// 创建日志目录
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(err)
	}

	// 配置日志输出文件
	hook := &lumberjack.Logger{
		Filename:   filepath.Join(logDir, "app.log"), // 日志文件路径
		MaxSize:    100,                              // 每个日志文件最大尺寸，单位 MB
		MaxBackups: 30,                               // 保留的旧文件最大数量
		MaxAge:     7,                                // 保留旧文件的最大天数
		Compress:   true,                             // 是否压缩旧文件
	}

	// 编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 创建核心
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                          // JSON 编码器
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook)), // 同时输出到控制台和文件
		zap.NewAtomicLevelAt(zap.InfoLevel),                                            // 日志级别
	)

	// 创建 logger
	log = zap.New(core,
		zap.AddCaller(),                   // 添加调用者信息
		zap.AddCallerSkip(1),              // 跳过一层调用栈
		zap.AddStacktrace(zap.ErrorLevel), // Error 级别及以上添加堆栈信息
	)
}

// Debug 输出 debug 级别日志
func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

// Info 输出 info 级别日志
func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

// Warn 输出 warn 级别日志
func Warn(msg string, fields ...zap.Field) {
	log.Warn(msg, fields...)
}

// Error 输出 error 级别日志
func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

// Fatal 输出 fatal 级别日志
func Fatal(msg string, fields ...zap.Field) {
	log.Fatal(msg, fields...)
}

// WithFields 创建带有固定字段的 logger
func WithFields(fields ...zap.Field) *zap.Logger {
	return log.With(fields...)
}

// Sync 同步日志缓冲
func Sync() {
	log.Sync()
}
