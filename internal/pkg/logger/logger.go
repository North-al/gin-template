package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

var log *slog.Logger

func InitLogger() {
	// 创建日志目录
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(err)
	}

	// 配置日志输出文件，按天和大小分割
	hook := &lumberjack.Logger{
		Filename:   filepath.Join(logDir, time.Now().Format("2006-01-02")+".log"), // 按天分割日志文件
		MaxSize:    100,                                                           // 每个日志文件最大尺寸，单位 MB
		MaxBackups: 60,                                                            // 保留的旧文件最大数量
		MaxAge:     60,                                                            // 保留旧文件的最大天数
		Compress:   true,                                                          // 是否压缩旧文件
		LocalTime:  true,                                                          // 使用本地时间
	}

	// 创建多输出 writer
	writer := io.MultiWriter(os.Stdout, hook)

	// 创建 JSON handler，增强日志级别显示
	handler := slog.NewJSONHandler(writer, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel := level.String()
				switch level {
				case slog.LevelDebug:
					levelLabel = "DEBUG"
				case slog.LevelInfo:
					levelLabel = "INFO"
				case slog.LevelWarn:
					levelLabel = "WARN"
				case slog.LevelError:
					levelLabel = "ERROR"
				}
				return slog.Attr{
					Key:   slog.LevelKey,
					Value: slog.StringValue(levelLabel),
				}
			}
			return a
		},
	})

	// 创建 logger
	log = slog.New(handler)
}

// Debug 输出 debug 级别日志
func Debug(msg string, args ...any) {
	log.Debug(msg, args...)
}

// Info 输出 info 级别日志
func Info(msg string, args ...any) {
	log.Info(msg, args...)
}

// Warn 输出 warn 级别日志
func Warn(msg string, args ...any) {
	log.Warn(msg, args...)
}

// Error 输出 error 级别日志
func Error(msg string, args ...any) {
	log.Error(msg, args...)
}

// Fatal 输出 fatal 级别日志
func Fatal(msg string, args ...any) {
	log.Error(msg, args...)
	os.Exit(1)
}

// WithFields 创建带有固定字段的 logger
func WithFields(args ...any) *slog.Logger {
	return log.With(args...)
}

// WithContext 从上下文中获取 logger
func WithContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(slog.Logger{}).(*slog.Logger); ok {
		return logger
	}
	return log
}

// NewContext 创建带有 logger 的上下文
func NewContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, slog.Logger{}, logger)
}
